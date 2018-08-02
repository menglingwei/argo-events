// +build !ignore_autogenerated

/*
Copyright 2018 The Argo Project Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Event) DeepCopyInto(out *Event) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Event.
func (in *Event) DeepCopy() *Event {
	if in == nil {
		return nil
	}
	out := new(Event)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Event) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventCondition) DeepCopyInto(out *EventCondition) {
	*out = *in
	if in.Sensor != nil {
		in, out := &in.Sensor, &out.Sensor
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Start != nil {
		in, out := &in.Start, &out.Start
		*out = (*in).DeepCopy()
	}
	if in.Stop != nil {
		in, out := &in.Stop, &out.Stop
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventCondition.
func (in *EventCondition) DeepCopy() *EventCondition {
	if in == nil {
		return nil
	}
	out := new(EventCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSpec) DeepCopyInto(out *EventSpec) {
	*out = *in
	if in.Deadline != nil {
		in, out := &in.Deadline, &out.Deadline
		*out = (*in).DeepCopy()
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]EventCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSpec.
func (in *EventSpec) DeepCopy() *EventSpec {
	if in == nil {
		return nil
	}
	out := new(EventSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventState) DeepCopyInto(out *EventState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventState.
func (in *EventState) DeepCopy() *EventState {
	if in == nil {
		return nil
	}
	out := new(EventState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventStatus) DeepCopyInto(out *EventStatus) {
	*out = *in
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]EventState, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventStatus.
func (in *EventStatus) DeepCopy() *EventStatus {
	if in == nil {
		return nil
	}
	out := new(EventStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Sensor) DeepCopyInto(out *Sensor) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Sensor.
func (in *Sensor) DeepCopy() *Sensor {
	if in == nil {
		return nil
	}
	out := new(Sensor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Sensor) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorList) DeepCopyInto(out *SensorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Sensor, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorList.
func (in *SensorList) DeepCopy() *SensorList {
	if in == nil {
		return nil
	}
	out := new(SensorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SensorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorSpec) DeepCopyInto(out *SensorSpec) {
	*out = *in
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = new(v1.Container)
		(*in).DeepCopyInto(*out)
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = new([]v1.Volume)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.Volume, len(*in))
			for i := range *in {
				(*in)[i].DeepCopyInto(&(*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorSpec.
func (in *SensorSpec) DeepCopy() *SensorSpec {
	if in == nil {
		return nil
	}
	out := new(SensorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorState) DeepCopyInto(out *SensorState) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorState.
func (in *SensorState) DeepCopy() *SensorState {
	if in == nil {
		return nil
	}
	out := new(SensorState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensorStatus) DeepCopyInto(out *SensorStatus) {
	*out = *in
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]SensorState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensorStatus.
func (in *SensorStatus) DeepCopy() *SensorStatus {
	if in == nil {
		return nil
	}
	out := new(SensorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Trigger) DeepCopyInto(out *Trigger) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Trigger.
func (in *Trigger) DeepCopy() *Trigger {
	if in == nil {
		return nil
	}
	out := new(Trigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Trigger) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerSpec) DeepCopyInto(out *TriggerSpec) {
	*out = *in
	if in.Event != nil {
		in, out := &in.Event, &out.Event
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.Executor != nil {
		in, out := &in.Executor, &out.Executor
		*out = new(v1.Container)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerSpec.
func (in *TriggerSpec) DeepCopy() *TriggerSpec {
	if in == nil {
		return nil
	}
	out := new(TriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerState) DeepCopyInto(out *TriggerState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerState.
func (in *TriggerState) DeepCopy() *TriggerState {
	if in == nil {
		return nil
	}
	out := new(TriggerState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TriggerStatus) DeepCopyInto(out *TriggerStatus) {
	*out = *in
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]TriggerState, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TriggerStatus.
func (in *TriggerStatus) DeepCopy() *TriggerStatus {
	if in == nil {
		return nil
	}
	out := new(TriggerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Wire) DeepCopyInto(out *Wire) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Wire.
func (in *Wire) DeepCopy() *Wire {
	if in == nil {
		return nil
	}
	out := new(Wire)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Wire) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireSpec) DeepCopyInto(out *WireSpec) {
	*out = *in
	in.Receiver.DeepCopyInto(&out.Receiver)
	in.Dispatcher.DeepCopyInto(&out.Dispatcher)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireSpec.
func (in *WireSpec) DeepCopy() *WireSpec {
	if in == nil {
		return nil
	}
	out := new(WireSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireState) DeepCopyInto(out *WireState) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireState.
func (in *WireState) DeepCopy() *WireState {
	if in == nil {
		return nil
	}
	out := new(WireState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WireStatus) DeepCopyInto(out *WireStatus) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	if in.States != nil {
		in, out := &in.States, &out.States
		*out = make([]WireState, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WireStatus.
func (in *WireStatus) DeepCopy() *WireStatus {
	if in == nil {
		return nil
	}
	out := new(WireStatus)
	in.DeepCopyInto(out)
	return out
}