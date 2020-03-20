/*
Copyright 2019 The Crossplane Authors.

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

// Code generated by angryjet. DO NOT EDIT.

package v1alpha1

import (
	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetBindingPhase of this MySQLInstance.
func (cm *MySQLInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetClassReference of this MySQLInstance.
func (cm *MySQLInstance) GetClassReference() *corev1.ObjectReference {
	return cm.Spec.ClassReference
}

// GetClassSelector of this MySQLInstance.
func (cm *MySQLInstance) GetClassSelector() *metav1.LabelSelector {
	return cm.Spec.ClassSelector
}

// GetCondition of this MySQLInstance.
func (cm *MySQLInstance) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return cm.Status.GetCondition(ct)
}

// GetResourceReference of this MySQLInstance.
func (cm *MySQLInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this MySQLInstance.
func (cm *MySQLInstance) GetWriteConnectionSecretToReference() *runtimev1alpha1.LocalSecretReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this MySQLInstance.
func (cm *MySQLInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetClassReference of this MySQLInstance.
func (cm *MySQLInstance) SetClassReference(r *corev1.ObjectReference) {
	cm.Spec.ClassReference = r
}

// SetClassSelector of this MySQLInstance.
func (cm *MySQLInstance) SetClassSelector(s *metav1.LabelSelector) {
	cm.Spec.ClassSelector = s
}

// SetConditions of this MySQLInstance.
func (cm *MySQLInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetResourceReference of this MySQLInstance.
func (cm *MySQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this MySQLInstance.
func (cm *MySQLInstance) SetWriteConnectionSecretToReference(r *runtimev1alpha1.LocalSecretReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this NoSQLInstance.
func (cm *NoSQLInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetClassReference of this NoSQLInstance.
func (cm *NoSQLInstance) GetClassReference() *corev1.ObjectReference {
	return cm.Spec.ClassReference
}

// GetClassSelector of this NoSQLInstance.
func (cm *NoSQLInstance) GetClassSelector() *metav1.LabelSelector {
	return cm.Spec.ClassSelector
}

// GetCondition of this NoSQLInstance.
func (cm *NoSQLInstance) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return cm.Status.GetCondition(ct)
}

// GetResourceReference of this NoSQLInstance.
func (cm *NoSQLInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this NoSQLInstance.
func (cm *NoSQLInstance) GetWriteConnectionSecretToReference() *runtimev1alpha1.LocalSecretReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this NoSQLInstance.
func (cm *NoSQLInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetClassReference of this NoSQLInstance.
func (cm *NoSQLInstance) SetClassReference(r *corev1.ObjectReference) {
	cm.Spec.ClassReference = r
}

// SetClassSelector of this NoSQLInstance.
func (cm *NoSQLInstance) SetClassSelector(s *metav1.LabelSelector) {
	cm.Spec.ClassSelector = s
}

// SetConditions of this NoSQLInstance.
func (cm *NoSQLInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetResourceReference of this NoSQLInstance.
func (cm *NoSQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this NoSQLInstance.
func (cm *NoSQLInstance) SetWriteConnectionSecretToReference(r *runtimev1alpha1.LocalSecretReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}

// GetBindingPhase of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetBindingPhase() runtimev1alpha1.BindingPhase {
	return cm.Status.GetBindingPhase()
}

// GetClassReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetClassReference() *corev1.ObjectReference {
	return cm.Spec.ClassReference
}

// GetClassSelector of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetClassSelector() *metav1.LabelSelector {
	return cm.Spec.ClassSelector
}

// GetCondition of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetCondition(ct runtimev1alpha1.ConditionType) runtimev1alpha1.Condition {
	return cm.Status.GetCondition(ct)
}

// GetResourceReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetResourceReference() *corev1.ObjectReference {
	return cm.Spec.ResourceReference
}

// GetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) GetWriteConnectionSecretToReference() *runtimev1alpha1.LocalSecretReference {
	return cm.Spec.WriteConnectionSecretToReference
}

// SetBindingPhase of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetBindingPhase(p runtimev1alpha1.BindingPhase) {
	cm.Status.SetBindingPhase(p)
}

// SetClassReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetClassReference(r *corev1.ObjectReference) {
	cm.Spec.ClassReference = r
}

// SetClassSelector of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetClassSelector(s *metav1.LabelSelector) {
	cm.Spec.ClassSelector = s
}

// SetConditions of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetConditions(c ...runtimev1alpha1.Condition) {
	cm.Status.SetConditions(c...)
}

// SetResourceReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetResourceReference(r *corev1.ObjectReference) {
	cm.Spec.ResourceReference = r
}

// SetWriteConnectionSecretToReference of this PostgreSQLInstance.
func (cm *PostgreSQLInstance) SetWriteConnectionSecretToReference(r *runtimev1alpha1.LocalSecretReference) {
	cm.Spec.WriteConnectionSecretToReference = r
}
