package a1aplha1

import (
"k8s.io/apimachinery/pkg/runtime/schema"
"k8s.io/apimachinery/pkg/runtime"
metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


var SchemeGroupVersion  = schema.GroupVersion{
    Group : "manansaini.dev",
    Version: "v1aplha1",
}


var (
    SchemeBuilder runtime.SchemeBuilder
)

func init() {
    SchemeBuilder.Register(addKnownTypes)
 }

 func addKnownTypes(schem *runtime.Scheme) error {
     scheme.AddKnownTypes(SchemeGroupVersion, &Kluster{}, &KlusterList{})


     metav1.AddToGroupVersion(scheme , SchemeGroupVersion)
    return nil
 }
