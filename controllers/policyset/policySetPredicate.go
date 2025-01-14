// Copyright (c) 2022 Red Hat, Inc.
// Copyright Contributors to the Open Cluster Management project

package controllers

import (
	"k8s.io/apimachinery/pkg/api/equality"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"

	policiesv1 "github.com/stolostron/governance-policy-propagator/api/v1"
)

var policySetPredicateFuncs = predicate.Funcs{
	UpdateFunc: func(e event.UpdateEvent) bool {
		// nolint: forcetypeassert
		policySetObjNew := e.ObjectNew.(*policiesv1.PolicySet)
		// nolint: forcetypeassert
		policySetObjOld := e.ObjectOld.(*policiesv1.PolicySet)

		return !equality.Semantic.DeepEqual(
			policySetObjNew.Spec.Policies,
			policySetObjOld.Spec.Policies,
		)
	},
	CreateFunc: func(e event.CreateEvent) bool {
		return true
	},
	DeleteFunc: func(e event.DeleteEvent) bool {
		return true
	},
}
