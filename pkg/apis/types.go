package apis

import (
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// ResourceInfo helps define a k8s resource
type ResourceInfo struct {
	group        string
	version      string
	resourceName string
	singularName string
	kind         string
	newObj       func() runtime.Object
	newList      func() runtime.Object
}

func NewResourceInfo(group, version, resourceName, singularName, kind string,
	newObj func() runtime.Object, newList func() runtime.Object) ResourceInfo {
	return ResourceInfo{group, version, resourceName, singularName, kind, newObj, newList}
}

func (info *ResourceInfo) GetSingularName() string {
	return info.singularName
}

// TypeMeta returns k8s type
func (info *ResourceInfo) TypeMeta() metav1.TypeMeta {
	return metav1.TypeMeta{
		Kind:       info.kind,
		APIVersion: info.group + "/" + info.version,
	}
}

func (info *ResourceInfo) GroupVersion() schema.GroupVersion {
	return schema.GroupVersion{
		Group:   info.group,
		Version: info.version,
	}
}

func (info *ResourceInfo) GroupResource() schema.GroupResource {
	return schema.GroupResource{
		Group:    info.group,
		Resource: info.resourceName,
	}
}

func (info *ResourceInfo) SingularGroupResource() schema.GroupResource {
	return schema.GroupResource{
		Group:    info.group,
		Resource: info.singularName,
	}
}

func (info *ResourceInfo) GroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    info.group,
		Version:  info.version,
		Resource: info.resourceName,
	}
}

func (info *ResourceInfo) StoragePath(sub ...string) string {
	switch len(sub) {
	case 0:
		return info.resourceName
	case 1:
		return info.resourceName + "/" + sub[0]
	}
	panic("invalid subresource path")
}

func (info *ResourceInfo) NewFunc() runtime.Object {
	return info.newObj()
}

func (info *ResourceInfo) NewListFunc() runtime.Object {
	return info.newList()
}

func (info *ResourceInfo) NewNotFound(name string) *errors.StatusError {
	return errors.NewNotFound(info.SingularGroupResource(), name)
}
