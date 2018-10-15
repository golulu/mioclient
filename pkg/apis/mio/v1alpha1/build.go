package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Foo is a specification for a Foo resource
type Build struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata.
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the desired behavior of the Deployment.
	// +optional
	Spec BuildSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Most recently observed status of the Deployment.
	// +optional
	Status BuildStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

type BuildStatus struct {
	King           string        `json:"king" protobuf:"bytes,1,opt,name=kind"`
	Name           string        `json:"name" protobuf:"bytes,2,opt,name=name"`
	Namespace      string        `json:"namespace" protobuf:"bytes,3,opt,name=namespace"`
	Phase          string        `json:"phase"  protobuf:"bytes,4,opt,name=phase"`
	Stages         []BuildStages `json:"stages" protobuf:"bytes,5,opt,name=stages"`
	StartTimestamp int64         `json:"startTimestamp" protobuf:"bytes,6,opt,name=startTimestamp"`
}

type BuildSpec struct {
	CloneConfig BuildCloneConfig `json:"cloneConfig"  protobuf:"bytes,1,opt,name=cloneConfig"`
	App         string           `json:"app"  protobuf:"bytes,1,opt,name=app"`
	//代码类型
	CodeType   string       `json:"codeType"  protobuf:"bytes,1,opt,name=codeType"`
	CompileCmd []CompileCmd `json:"compileCmd"  protobuf:"bytes,1,opt,name=compileCmd"`
	//获取类型
	CloneType string `json:"cloneType"  protobuf:"bytes,1,opt,name=cloneType"`
	//基础镜像包
	S2iImage string `json:"s2iImage"  protobuf:"bytes,1,opt,name=s2iImage"`
	//版本
	Tags       []string `json:"tags"  protobuf:"bytes,1,opt,name=tags"`
	DockerFile []string `json:"dockerFile"  protobuf:"bytes,1,opt,name=dockerFile"`
}

type BuildCloneConfig struct {
	Url      string `json:"url"  protobuf:"bytes,1,opt,name=url"`
	Branch   string `json:"branch"  protobuf:"bytes,1,opt,name=branch"`
	DstDir   string `json:"dstDir"  protobuf:"bytes,1,opt,name=dstDir"`
	Username string `json:"username"  protobuf:"bytes,1,opt,name=username"`
	Password string `json:"password"  protobuf:"bytes,1,opt,name=password"`
}
type CompileType string

const (
	Command CompileType = "command"
	Script  CompileType = "script"
)

type CompileCmd struct {
	Type          CompileType `json:"CompileType"  protobuf:"bytes,1,opt,name=compileType"`
	Script        string      `json:"Script"  protobuf:"bytes,1,opt,name=script"`
	CommandName   string      `json:"commandName"  protobuf:"bytes,1,opt,name=commandName"`
	CommandParams []string    `json:"params"  protobuf:"bytes,1,opt,name=params"`
}

type BuildStages struct {
	Name                 string `json:"name" protobuf:"bytes,1,opt,name=name"`
	StartTime            int64  `json:"startTime" protobuf:"bytes,2,opt,name=startTime"`
	DurationMilliseconds int64  `json:"durationMilliseconds" protobuf:"bytes,3,opt,name=durationMilliseconds"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// FooList is a list of Foo resources
type BuildList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []Build `json:"items"`
}
