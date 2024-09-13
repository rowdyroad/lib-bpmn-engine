package bpmn_engine

import (
	"github.com/nitram509/lib-bpmn-engine/pkg/bpmn_engine/exporter"
	"github.com/nitram509/lib-bpmn-engine/pkg/spec/BPMN20"
)

// AddEventExporter registers an EventExporter instance
func (state *BpmnEngineState) AddEventExporter(exporter exporter.EventExporter) {
	state.exporters = append(state.exporters, exporter)
}

func (state *BpmnEngineState) exportNewProcessEvent(processInfo ProcessInfo, xmlData []byte, resourceName string, checksum string) {
	event := exporter.ProcessEvent{
		ProcessId:    processInfo.BpmnProcessId,
		ProcessKey:   processInfo.ProcessKey,
		Version:      processInfo.Version,
		XmlData:      xmlData,
		ResourceName: resourceName,
		Checksum:     checksum,
	}
	for _, exp := range state.exporters {
		exp.NewProcessEvent(&event)
	}
}

func (state *BpmnEngineState) exportEndProcessEvent(process ProcessInfo, processInstance processInstanceInfo) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.InstanceKey,
	}
	for _, exp := range state.exporters {
		exp.EndProcessEvent(&event)
	}
}

func (state *BpmnEngineState) exportProcessInstanceEvent(process ProcessInfo, processInstance processInstanceInfo) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.InstanceKey,
	}
	for _, exp := range state.exporters {
		exp.NewProcessInstanceEvent(&event)
	}
}

func (state *BpmnEngineState) exportElementEvent(process ProcessInfo, processInstance processInstanceInfo, element BPMN20.BaseElement, intent exporter.Intent) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.InstanceKey,
	}
	info := exporter.ElementInfo{
		BpmnElementType: string(element.GetType()),
		ElementId:       element.GetId(),
		Intent:          string(intent),
	}
	for _, exp := range state.exporters {
		exp.NewElementEvent(&event, &info)
	}
}

func (state *BpmnEngineState) exportSequenceFlowEvent(process ProcessInfo, processInstance processInstanceInfo, flow BPMN20.TSequenceFlow) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.BpmnProcessId,
		ProcessKey:         process.ProcessKey,
		Version:            process.Version,
		ProcessInstanceKey: processInstance.InstanceKey,
	}
	info := exporter.ElementInfo{
		BpmnElementType: string(BPMN20.SequenceFlow),
		ElementId:       flow.Id,
		Intent:          string(exporter.SequenceFlowTaken),
	}
	for _, exp := range state.exporters {
		exp.NewElementEvent(&event, &info)
	}
}

func (state *BpmnEngineState) exportRemoveProcessEvent(process *ProcessInfo) {
	for _, exp := range state.exporters {
		exp.RemoveProcessEvent(&exporter.ProcessInstanceEvent{
			ProcessId:  process.BpmnProcessId,
			ProcessKey: process.ProcessKey,
			Version:    process.Version,
		})
	}
}

func (state *BpmnEngineState) exportRemoveProcessInstanceEvent(process *processInstanceInfo) {
	for _, exp := range state.exporters {
		exp.RemoveProcessInstanceEvent(&exporter.ProcessInstanceEvent{
			ProcessId:          process.ProcessInfo.BpmnProcessId,
			ProcessKey:         process.ProcessInfo.ProcessKey,
			Version:            process.ProcessInfo.Version,
			ProcessInstanceKey: process.InstanceKey,
		})
	}
}
func (state *BpmnEngineState) exportRemoveMessageSubscriptionEvent(process *processInstanceInfo, message *MessageSubscription) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.ProcessInfo.BpmnProcessId,
		ProcessKey:         process.ProcessInfo.ProcessKey,
		Version:            process.ProcessInfo.Version,
		ProcessInstanceKey: process.InstanceKey,
	}
	info := exporter.ElementInfo{
		BpmnElementType: message.Name,
		ElementId:       message.ElementId,
		Intent:          string(message.MessageState),
		InstanceKey:     message.ElementInstanceKey,
	}
	for _, exp := range state.exporters {
		exp.RemoveMessageSubscriptionEvent(&event, &info)
	}
}
func (state *BpmnEngineState) exportRemoveJobEvent(process *processInstanceInfo, job *job) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.ProcessInfo.BpmnProcessId,
		ProcessKey:         process.ProcessInfo.ProcessKey,
		Version:            process.ProcessInfo.Version,
		ProcessInstanceKey: process.InstanceKey,
	}
	info := exporter.ElementInfo{
		ElementId:   job.ElementId,
		Intent:      string(job.JobState),
		InstanceKey: job.ElementInstanceKey,
		Key:         job.JobKey,
	}
	for _, exp := range state.exporters {
		exp.RemoveJobEvent(&event, &info)
	}
}
func (state *BpmnEngineState) exportRemoveTimerEvent(process *processInstanceInfo, timer *Timer) {
	event := exporter.ProcessInstanceEvent{
		ProcessId:          process.ProcessInfo.BpmnProcessId,
		ProcessKey:         process.ProcessInfo.ProcessKey,
		Version:            process.ProcessInfo.Version,
		ProcessInstanceKey: process.InstanceKey,
	}
	info := exporter.ElementInfo{
		ElementId:   timer.ElementId,
		Intent:      string(timer.TimerState),
		InstanceKey: timer.ElementInstanceKey,
	}
	for _, exp := range state.exporters {
		exp.RemoveTimerEvent(&event, &info)
	}
}
