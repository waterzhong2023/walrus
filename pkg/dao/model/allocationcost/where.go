// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package allocationcost

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/types"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldID, id))
}

// StartTime applies equality check predicate on the "startTime" field. It's identical to StartTimeEQ.
func StartTime(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldStartTime, v))
}

// EndTime applies equality check predicate on the "endTime" field. It's identical to EndTimeEQ.
func EndTime(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldEndTime, v))
}

// Minutes applies equality check predicate on the "minutes" field. It's identical to MinutesEQ.
func Minutes(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldMinutes, v))
}

// ConnectorID applies equality check predicate on the "connectorID" field. It's identical to ConnectorIDEQ.
func ConnectorID(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldConnectorID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldName, v))
}

// Fingerprint applies equality check predicate on the "fingerprint" field. It's identical to FingerprintEQ.
func Fingerprint(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldFingerprint, v))
}

// ClusterName applies equality check predicate on the "clusterName" field. It's identical to ClusterNameEQ.
func ClusterName(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldClusterName, v))
}

// Namespace applies equality check predicate on the "namespace" field. It's identical to NamespaceEQ.
func Namespace(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldNamespace, v))
}

// Node applies equality check predicate on the "node" field. It's identical to NodeEQ.
func Node(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldNode, v))
}

// Controller applies equality check predicate on the "controller" field. It's identical to ControllerEQ.
func Controller(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldController, v))
}

// ControllerKind applies equality check predicate on the "controllerKind" field. It's identical to ControllerKindEQ.
func ControllerKind(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldControllerKind, v))
}

// Pod applies equality check predicate on the "pod" field. It's identical to PodEQ.
func Pod(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPod, v))
}

// Container applies equality check predicate on the "container" field. It's identical to ContainerEQ.
func Container(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldContainer, v))
}

// TotalCost applies equality check predicate on the "totalCost" field. It's identical to TotalCostEQ.
func TotalCost(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldTotalCost, v))
}

// Currency applies equality check predicate on the "currency" field. It's identical to CurrencyEQ.
func Currency(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCurrency, v))
}

// CpuCost applies equality check predicate on the "cpuCost" field. It's identical to CpuCostEQ.
func CpuCost(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCost, v))
}

// CpuCoreRequest applies equality check predicate on the "cpuCoreRequest" field. It's identical to CpuCoreRequestEQ.
func CpuCoreRequest(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreRequest, v))
}

// GpuCost applies equality check predicate on the "gpuCost" field. It's identical to GpuCostEQ.
func GpuCost(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldGpuCost, v))
}

// GpuCount applies equality check predicate on the "gpuCount" field. It's identical to GpuCountEQ.
func GpuCount(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldGpuCount, v))
}

// RamCost applies equality check predicate on the "ramCost" field. It's identical to RamCostEQ.
func RamCost(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamCost, v))
}

// RamByteRequest applies equality check predicate on the "ramByteRequest" field. It's identical to RamByteRequestEQ.
func RamByteRequest(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteRequest, v))
}

// PvCost applies equality check predicate on the "pvCost" field. It's identical to PvCostEQ.
func PvCost(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPvCost, v))
}

// PvBytes applies equality check predicate on the "pvBytes" field. It's identical to PvBytesEQ.
func PvBytes(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPvBytes, v))
}

// CpuCoreUsageAverage applies equality check predicate on the "cpuCoreUsageAverage" field. It's identical to CpuCoreUsageAverageEQ.
func CpuCoreUsageAverage(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageMax applies equality check predicate on the "cpuCoreUsageMax" field. It's identical to CpuCoreUsageMaxEQ.
func CpuCoreUsageMax(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreUsageMax, v))
}

// RamByteUsageAverage applies equality check predicate on the "ramByteUsageAverage" field. It's identical to RamByteUsageAverageEQ.
func RamByteUsageAverage(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteUsageAverage, v))
}

// RamByteUsageMax applies equality check predicate on the "ramByteUsageMax" field. It's identical to RamByteUsageMaxEQ.
func RamByteUsageMax(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteUsageMax, v))
}

// StartTimeEQ applies the EQ predicate on the "startTime" field.
func StartTimeEQ(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldStartTime, v))
}

// StartTimeNEQ applies the NEQ predicate on the "startTime" field.
func StartTimeNEQ(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldStartTime, v))
}

// StartTimeIn applies the In predicate on the "startTime" field.
func StartTimeIn(vs ...time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldStartTime, vs...))
}

// StartTimeNotIn applies the NotIn predicate on the "startTime" field.
func StartTimeNotIn(vs ...time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldStartTime, vs...))
}

// StartTimeGT applies the GT predicate on the "startTime" field.
func StartTimeGT(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldStartTime, v))
}

// StartTimeGTE applies the GTE predicate on the "startTime" field.
func StartTimeGTE(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldStartTime, v))
}

// StartTimeLT applies the LT predicate on the "startTime" field.
func StartTimeLT(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldStartTime, v))
}

// StartTimeLTE applies the LTE predicate on the "startTime" field.
func StartTimeLTE(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldStartTime, v))
}

// EndTimeEQ applies the EQ predicate on the "endTime" field.
func EndTimeEQ(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldEndTime, v))
}

// EndTimeNEQ applies the NEQ predicate on the "endTime" field.
func EndTimeNEQ(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldEndTime, v))
}

// EndTimeIn applies the In predicate on the "endTime" field.
func EndTimeIn(vs ...time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldEndTime, vs...))
}

// EndTimeNotIn applies the NotIn predicate on the "endTime" field.
func EndTimeNotIn(vs ...time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldEndTime, vs...))
}

// EndTimeGT applies the GT predicate on the "endTime" field.
func EndTimeGT(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldEndTime, v))
}

// EndTimeGTE applies the GTE predicate on the "endTime" field.
func EndTimeGTE(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldEndTime, v))
}

// EndTimeLT applies the LT predicate on the "endTime" field.
func EndTimeLT(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldEndTime, v))
}

// EndTimeLTE applies the LTE predicate on the "endTime" field.
func EndTimeLTE(v time.Time) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldEndTime, v))
}

// MinutesEQ applies the EQ predicate on the "minutes" field.
func MinutesEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldMinutes, v))
}

// MinutesNEQ applies the NEQ predicate on the "minutes" field.
func MinutesNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldMinutes, v))
}

// MinutesIn applies the In predicate on the "minutes" field.
func MinutesIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldMinutes, vs...))
}

// MinutesNotIn applies the NotIn predicate on the "minutes" field.
func MinutesNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldMinutes, vs...))
}

// MinutesGT applies the GT predicate on the "minutes" field.
func MinutesGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldMinutes, v))
}

// MinutesGTE applies the GTE predicate on the "minutes" field.
func MinutesGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldMinutes, v))
}

// MinutesLT applies the LT predicate on the "minutes" field.
func MinutesLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldMinutes, v))
}

// MinutesLTE applies the LTE predicate on the "minutes" field.
func MinutesLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldMinutes, v))
}

// ConnectorIDEQ applies the EQ predicate on the "connectorID" field.
func ConnectorIDEQ(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldConnectorID, v))
}

// ConnectorIDNEQ applies the NEQ predicate on the "connectorID" field.
func ConnectorIDNEQ(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldConnectorID, v))
}

// ConnectorIDIn applies the In predicate on the "connectorID" field.
func ConnectorIDIn(vs ...types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldConnectorID, vs...))
}

// ConnectorIDNotIn applies the NotIn predicate on the "connectorID" field.
func ConnectorIDNotIn(vs ...types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldConnectorID, vs...))
}

// ConnectorIDGT applies the GT predicate on the "connectorID" field.
func ConnectorIDGT(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldConnectorID, v))
}

// ConnectorIDGTE applies the GTE predicate on the "connectorID" field.
func ConnectorIDGTE(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldConnectorID, v))
}

// ConnectorIDLT applies the LT predicate on the "connectorID" field.
func ConnectorIDLT(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldConnectorID, v))
}

// ConnectorIDLTE applies the LTE predicate on the "connectorID" field.
func ConnectorIDLTE(v types.ID) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldConnectorID, v))
}

// ConnectorIDContains applies the Contains predicate on the "connectorID" field.
func ConnectorIDContains(v types.ID) predicate.AllocationCost {
	vc := string(v)
	return predicate.AllocationCost(sql.FieldContains(FieldConnectorID, vc))
}

// ConnectorIDHasPrefix applies the HasPrefix predicate on the "connectorID" field.
func ConnectorIDHasPrefix(v types.ID) predicate.AllocationCost {
	vc := string(v)
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldConnectorID, vc))
}

// ConnectorIDHasSuffix applies the HasSuffix predicate on the "connectorID" field.
func ConnectorIDHasSuffix(v types.ID) predicate.AllocationCost {
	vc := string(v)
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldConnectorID, vc))
}

// ConnectorIDEqualFold applies the EqualFold predicate on the "connectorID" field.
func ConnectorIDEqualFold(v types.ID) predicate.AllocationCost {
	vc := string(v)
	return predicate.AllocationCost(sql.FieldEqualFold(FieldConnectorID, vc))
}

// ConnectorIDContainsFold applies the ContainsFold predicate on the "connectorID" field.
func ConnectorIDContainsFold(v types.ID) predicate.AllocationCost {
	vc := string(v)
	return predicate.AllocationCost(sql.FieldContainsFold(FieldConnectorID, vc))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldName, v))
}

// FingerprintEQ applies the EQ predicate on the "fingerprint" field.
func FingerprintEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldFingerprint, v))
}

// FingerprintNEQ applies the NEQ predicate on the "fingerprint" field.
func FingerprintNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldFingerprint, v))
}

// FingerprintIn applies the In predicate on the "fingerprint" field.
func FingerprintIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldFingerprint, vs...))
}

// FingerprintNotIn applies the NotIn predicate on the "fingerprint" field.
func FingerprintNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldFingerprint, vs...))
}

// FingerprintGT applies the GT predicate on the "fingerprint" field.
func FingerprintGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldFingerprint, v))
}

// FingerprintGTE applies the GTE predicate on the "fingerprint" field.
func FingerprintGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldFingerprint, v))
}

// FingerprintLT applies the LT predicate on the "fingerprint" field.
func FingerprintLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldFingerprint, v))
}

// FingerprintLTE applies the LTE predicate on the "fingerprint" field.
func FingerprintLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldFingerprint, v))
}

// FingerprintContains applies the Contains predicate on the "fingerprint" field.
func FingerprintContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldFingerprint, v))
}

// FingerprintHasPrefix applies the HasPrefix predicate on the "fingerprint" field.
func FingerprintHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldFingerprint, v))
}

// FingerprintHasSuffix applies the HasSuffix predicate on the "fingerprint" field.
func FingerprintHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldFingerprint, v))
}

// FingerprintEqualFold applies the EqualFold predicate on the "fingerprint" field.
func FingerprintEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldFingerprint, v))
}

// FingerprintContainsFold applies the ContainsFold predicate on the "fingerprint" field.
func FingerprintContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldFingerprint, v))
}

// ClusterNameEQ applies the EQ predicate on the "clusterName" field.
func ClusterNameEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldClusterName, v))
}

// ClusterNameNEQ applies the NEQ predicate on the "clusterName" field.
func ClusterNameNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldClusterName, v))
}

// ClusterNameIn applies the In predicate on the "clusterName" field.
func ClusterNameIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldClusterName, vs...))
}

// ClusterNameNotIn applies the NotIn predicate on the "clusterName" field.
func ClusterNameNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldClusterName, vs...))
}

// ClusterNameGT applies the GT predicate on the "clusterName" field.
func ClusterNameGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldClusterName, v))
}

// ClusterNameGTE applies the GTE predicate on the "clusterName" field.
func ClusterNameGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldClusterName, v))
}

// ClusterNameLT applies the LT predicate on the "clusterName" field.
func ClusterNameLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldClusterName, v))
}

// ClusterNameLTE applies the LTE predicate on the "clusterName" field.
func ClusterNameLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldClusterName, v))
}

// ClusterNameContains applies the Contains predicate on the "clusterName" field.
func ClusterNameContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldClusterName, v))
}

// ClusterNameHasPrefix applies the HasPrefix predicate on the "clusterName" field.
func ClusterNameHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldClusterName, v))
}

// ClusterNameHasSuffix applies the HasSuffix predicate on the "clusterName" field.
func ClusterNameHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldClusterName, v))
}

// ClusterNameIsNil applies the IsNil predicate on the "clusterName" field.
func ClusterNameIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldClusterName))
}

// ClusterNameNotNil applies the NotNil predicate on the "clusterName" field.
func ClusterNameNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldClusterName))
}

// ClusterNameEqualFold applies the EqualFold predicate on the "clusterName" field.
func ClusterNameEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldClusterName, v))
}

// ClusterNameContainsFold applies the ContainsFold predicate on the "clusterName" field.
func ClusterNameContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldClusterName, v))
}

// NamespaceEQ applies the EQ predicate on the "namespace" field.
func NamespaceEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldNamespace, v))
}

// NamespaceNEQ applies the NEQ predicate on the "namespace" field.
func NamespaceNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldNamespace, v))
}

// NamespaceIn applies the In predicate on the "namespace" field.
func NamespaceIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldNamespace, vs...))
}

// NamespaceNotIn applies the NotIn predicate on the "namespace" field.
func NamespaceNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldNamespace, vs...))
}

// NamespaceGT applies the GT predicate on the "namespace" field.
func NamespaceGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldNamespace, v))
}

// NamespaceGTE applies the GTE predicate on the "namespace" field.
func NamespaceGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldNamespace, v))
}

// NamespaceLT applies the LT predicate on the "namespace" field.
func NamespaceLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldNamespace, v))
}

// NamespaceLTE applies the LTE predicate on the "namespace" field.
func NamespaceLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldNamespace, v))
}

// NamespaceContains applies the Contains predicate on the "namespace" field.
func NamespaceContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldNamespace, v))
}

// NamespaceHasPrefix applies the HasPrefix predicate on the "namespace" field.
func NamespaceHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldNamespace, v))
}

// NamespaceHasSuffix applies the HasSuffix predicate on the "namespace" field.
func NamespaceHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldNamespace, v))
}

// NamespaceIsNil applies the IsNil predicate on the "namespace" field.
func NamespaceIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldNamespace))
}

// NamespaceNotNil applies the NotNil predicate on the "namespace" field.
func NamespaceNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldNamespace))
}

// NamespaceEqualFold applies the EqualFold predicate on the "namespace" field.
func NamespaceEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldNamespace, v))
}

// NamespaceContainsFold applies the ContainsFold predicate on the "namespace" field.
func NamespaceContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldNamespace, v))
}

// NodeEQ applies the EQ predicate on the "node" field.
func NodeEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldNode, v))
}

// NodeNEQ applies the NEQ predicate on the "node" field.
func NodeNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldNode, v))
}

// NodeIn applies the In predicate on the "node" field.
func NodeIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldNode, vs...))
}

// NodeNotIn applies the NotIn predicate on the "node" field.
func NodeNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldNode, vs...))
}

// NodeGT applies the GT predicate on the "node" field.
func NodeGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldNode, v))
}

// NodeGTE applies the GTE predicate on the "node" field.
func NodeGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldNode, v))
}

// NodeLT applies the LT predicate on the "node" field.
func NodeLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldNode, v))
}

// NodeLTE applies the LTE predicate on the "node" field.
func NodeLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldNode, v))
}

// NodeContains applies the Contains predicate on the "node" field.
func NodeContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldNode, v))
}

// NodeHasPrefix applies the HasPrefix predicate on the "node" field.
func NodeHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldNode, v))
}

// NodeHasSuffix applies the HasSuffix predicate on the "node" field.
func NodeHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldNode, v))
}

// NodeIsNil applies the IsNil predicate on the "node" field.
func NodeIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldNode))
}

// NodeNotNil applies the NotNil predicate on the "node" field.
func NodeNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldNode))
}

// NodeEqualFold applies the EqualFold predicate on the "node" field.
func NodeEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldNode, v))
}

// NodeContainsFold applies the ContainsFold predicate on the "node" field.
func NodeContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldNode, v))
}

// ControllerEQ applies the EQ predicate on the "controller" field.
func ControllerEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldController, v))
}

// ControllerNEQ applies the NEQ predicate on the "controller" field.
func ControllerNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldController, v))
}

// ControllerIn applies the In predicate on the "controller" field.
func ControllerIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldController, vs...))
}

// ControllerNotIn applies the NotIn predicate on the "controller" field.
func ControllerNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldController, vs...))
}

// ControllerGT applies the GT predicate on the "controller" field.
func ControllerGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldController, v))
}

// ControllerGTE applies the GTE predicate on the "controller" field.
func ControllerGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldController, v))
}

// ControllerLT applies the LT predicate on the "controller" field.
func ControllerLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldController, v))
}

// ControllerLTE applies the LTE predicate on the "controller" field.
func ControllerLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldController, v))
}

// ControllerContains applies the Contains predicate on the "controller" field.
func ControllerContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldController, v))
}

// ControllerHasPrefix applies the HasPrefix predicate on the "controller" field.
func ControllerHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldController, v))
}

// ControllerHasSuffix applies the HasSuffix predicate on the "controller" field.
func ControllerHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldController, v))
}

// ControllerIsNil applies the IsNil predicate on the "controller" field.
func ControllerIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldController))
}

// ControllerNotNil applies the NotNil predicate on the "controller" field.
func ControllerNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldController))
}

// ControllerEqualFold applies the EqualFold predicate on the "controller" field.
func ControllerEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldController, v))
}

// ControllerContainsFold applies the ContainsFold predicate on the "controller" field.
func ControllerContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldController, v))
}

// ControllerKindEQ applies the EQ predicate on the "controllerKind" field.
func ControllerKindEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldControllerKind, v))
}

// ControllerKindNEQ applies the NEQ predicate on the "controllerKind" field.
func ControllerKindNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldControllerKind, v))
}

// ControllerKindIn applies the In predicate on the "controllerKind" field.
func ControllerKindIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldControllerKind, vs...))
}

// ControllerKindNotIn applies the NotIn predicate on the "controllerKind" field.
func ControllerKindNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldControllerKind, vs...))
}

// ControllerKindGT applies the GT predicate on the "controllerKind" field.
func ControllerKindGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldControllerKind, v))
}

// ControllerKindGTE applies the GTE predicate on the "controllerKind" field.
func ControllerKindGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldControllerKind, v))
}

// ControllerKindLT applies the LT predicate on the "controllerKind" field.
func ControllerKindLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldControllerKind, v))
}

// ControllerKindLTE applies the LTE predicate on the "controllerKind" field.
func ControllerKindLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldControllerKind, v))
}

// ControllerKindContains applies the Contains predicate on the "controllerKind" field.
func ControllerKindContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldControllerKind, v))
}

// ControllerKindHasPrefix applies the HasPrefix predicate on the "controllerKind" field.
func ControllerKindHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldControllerKind, v))
}

// ControllerKindHasSuffix applies the HasSuffix predicate on the "controllerKind" field.
func ControllerKindHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldControllerKind, v))
}

// ControllerKindIsNil applies the IsNil predicate on the "controllerKind" field.
func ControllerKindIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldControllerKind))
}

// ControllerKindNotNil applies the NotNil predicate on the "controllerKind" field.
func ControllerKindNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldControllerKind))
}

// ControllerKindEqualFold applies the EqualFold predicate on the "controllerKind" field.
func ControllerKindEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldControllerKind, v))
}

// ControllerKindContainsFold applies the ContainsFold predicate on the "controllerKind" field.
func ControllerKindContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldControllerKind, v))
}

// PodEQ applies the EQ predicate on the "pod" field.
func PodEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPod, v))
}

// PodNEQ applies the NEQ predicate on the "pod" field.
func PodNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldPod, v))
}

// PodIn applies the In predicate on the "pod" field.
func PodIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldPod, vs...))
}

// PodNotIn applies the NotIn predicate on the "pod" field.
func PodNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldPod, vs...))
}

// PodGT applies the GT predicate on the "pod" field.
func PodGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldPod, v))
}

// PodGTE applies the GTE predicate on the "pod" field.
func PodGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldPod, v))
}

// PodLT applies the LT predicate on the "pod" field.
func PodLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldPod, v))
}

// PodLTE applies the LTE predicate on the "pod" field.
func PodLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldPod, v))
}

// PodContains applies the Contains predicate on the "pod" field.
func PodContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldPod, v))
}

// PodHasPrefix applies the HasPrefix predicate on the "pod" field.
func PodHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldPod, v))
}

// PodHasSuffix applies the HasSuffix predicate on the "pod" field.
func PodHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldPod, v))
}

// PodIsNil applies the IsNil predicate on the "pod" field.
func PodIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldPod))
}

// PodNotNil applies the NotNil predicate on the "pod" field.
func PodNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldPod))
}

// PodEqualFold applies the EqualFold predicate on the "pod" field.
func PodEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldPod, v))
}

// PodContainsFold applies the ContainsFold predicate on the "pod" field.
func PodContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldPod, v))
}

// ContainerEQ applies the EQ predicate on the "container" field.
func ContainerEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldContainer, v))
}

// ContainerNEQ applies the NEQ predicate on the "container" field.
func ContainerNEQ(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldContainer, v))
}

// ContainerIn applies the In predicate on the "container" field.
func ContainerIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldContainer, vs...))
}

// ContainerNotIn applies the NotIn predicate on the "container" field.
func ContainerNotIn(vs ...string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldContainer, vs...))
}

// ContainerGT applies the GT predicate on the "container" field.
func ContainerGT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldContainer, v))
}

// ContainerGTE applies the GTE predicate on the "container" field.
func ContainerGTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldContainer, v))
}

// ContainerLT applies the LT predicate on the "container" field.
func ContainerLT(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldContainer, v))
}

// ContainerLTE applies the LTE predicate on the "container" field.
func ContainerLTE(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldContainer, v))
}

// ContainerContains applies the Contains predicate on the "container" field.
func ContainerContains(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContains(FieldContainer, v))
}

// ContainerHasPrefix applies the HasPrefix predicate on the "container" field.
func ContainerHasPrefix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasPrefix(FieldContainer, v))
}

// ContainerHasSuffix applies the HasSuffix predicate on the "container" field.
func ContainerHasSuffix(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldHasSuffix(FieldContainer, v))
}

// ContainerIsNil applies the IsNil predicate on the "container" field.
func ContainerIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldContainer))
}

// ContainerNotNil applies the NotNil predicate on the "container" field.
func ContainerNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldContainer))
}

// ContainerEqualFold applies the EqualFold predicate on the "container" field.
func ContainerEqualFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEqualFold(FieldContainer, v))
}

// ContainerContainsFold applies the ContainsFold predicate on the "container" field.
func ContainerContainsFold(v string) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldContainsFold(FieldContainer, v))
}

// TotalCostEQ applies the EQ predicate on the "totalCost" field.
func TotalCostEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldTotalCost, v))
}

// TotalCostNEQ applies the NEQ predicate on the "totalCost" field.
func TotalCostNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldTotalCost, v))
}

// TotalCostIn applies the In predicate on the "totalCost" field.
func TotalCostIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldTotalCost, vs...))
}

// TotalCostNotIn applies the NotIn predicate on the "totalCost" field.
func TotalCostNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldTotalCost, vs...))
}

// TotalCostGT applies the GT predicate on the "totalCost" field.
func TotalCostGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldTotalCost, v))
}

// TotalCostGTE applies the GTE predicate on the "totalCost" field.
func TotalCostGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldTotalCost, v))
}

// TotalCostLT applies the LT predicate on the "totalCost" field.
func TotalCostLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldTotalCost, v))
}

// TotalCostLTE applies the LTE predicate on the "totalCost" field.
func TotalCostLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldTotalCost, v))
}

// CurrencyEQ applies the EQ predicate on the "currency" field.
func CurrencyEQ(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCurrency, v))
}

// CurrencyNEQ applies the NEQ predicate on the "currency" field.
func CurrencyNEQ(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldCurrency, v))
}

// CurrencyIn applies the In predicate on the "currency" field.
func CurrencyIn(vs ...int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldCurrency, vs...))
}

// CurrencyNotIn applies the NotIn predicate on the "currency" field.
func CurrencyNotIn(vs ...int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldCurrency, vs...))
}

// CurrencyGT applies the GT predicate on the "currency" field.
func CurrencyGT(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldCurrency, v))
}

// CurrencyGTE applies the GTE predicate on the "currency" field.
func CurrencyGTE(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldCurrency, v))
}

// CurrencyLT applies the LT predicate on the "currency" field.
func CurrencyLT(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldCurrency, v))
}

// CurrencyLTE applies the LTE predicate on the "currency" field.
func CurrencyLTE(v int) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldCurrency, v))
}

// CurrencyIsNil applies the IsNil predicate on the "currency" field.
func CurrencyIsNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIsNull(FieldCurrency))
}

// CurrencyNotNil applies the NotNil predicate on the "currency" field.
func CurrencyNotNil() predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotNull(FieldCurrency))
}

// CpuCostEQ applies the EQ predicate on the "cpuCost" field.
func CpuCostEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCost, v))
}

// CpuCostNEQ applies the NEQ predicate on the "cpuCost" field.
func CpuCostNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldCpuCost, v))
}

// CpuCostIn applies the In predicate on the "cpuCost" field.
func CpuCostIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldCpuCost, vs...))
}

// CpuCostNotIn applies the NotIn predicate on the "cpuCost" field.
func CpuCostNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldCpuCost, vs...))
}

// CpuCostGT applies the GT predicate on the "cpuCost" field.
func CpuCostGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldCpuCost, v))
}

// CpuCostGTE applies the GTE predicate on the "cpuCost" field.
func CpuCostGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldCpuCost, v))
}

// CpuCostLT applies the LT predicate on the "cpuCost" field.
func CpuCostLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldCpuCost, v))
}

// CpuCostLTE applies the LTE predicate on the "cpuCost" field.
func CpuCostLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldCpuCost, v))
}

// CpuCoreRequestEQ applies the EQ predicate on the "cpuCoreRequest" field.
func CpuCoreRequestEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreRequest, v))
}

// CpuCoreRequestNEQ applies the NEQ predicate on the "cpuCoreRequest" field.
func CpuCoreRequestNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldCpuCoreRequest, v))
}

// CpuCoreRequestIn applies the In predicate on the "cpuCoreRequest" field.
func CpuCoreRequestIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldCpuCoreRequest, vs...))
}

// CpuCoreRequestNotIn applies the NotIn predicate on the "cpuCoreRequest" field.
func CpuCoreRequestNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldCpuCoreRequest, vs...))
}

// CpuCoreRequestGT applies the GT predicate on the "cpuCoreRequest" field.
func CpuCoreRequestGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldCpuCoreRequest, v))
}

// CpuCoreRequestGTE applies the GTE predicate on the "cpuCoreRequest" field.
func CpuCoreRequestGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldCpuCoreRequest, v))
}

// CpuCoreRequestLT applies the LT predicate on the "cpuCoreRequest" field.
func CpuCoreRequestLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldCpuCoreRequest, v))
}

// CpuCoreRequestLTE applies the LTE predicate on the "cpuCoreRequest" field.
func CpuCoreRequestLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldCpuCoreRequest, v))
}

// GpuCostEQ applies the EQ predicate on the "gpuCost" field.
func GpuCostEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldGpuCost, v))
}

// GpuCostNEQ applies the NEQ predicate on the "gpuCost" field.
func GpuCostNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldGpuCost, v))
}

// GpuCostIn applies the In predicate on the "gpuCost" field.
func GpuCostIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldGpuCost, vs...))
}

// GpuCostNotIn applies the NotIn predicate on the "gpuCost" field.
func GpuCostNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldGpuCost, vs...))
}

// GpuCostGT applies the GT predicate on the "gpuCost" field.
func GpuCostGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldGpuCost, v))
}

// GpuCostGTE applies the GTE predicate on the "gpuCost" field.
func GpuCostGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldGpuCost, v))
}

// GpuCostLT applies the LT predicate on the "gpuCost" field.
func GpuCostLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldGpuCost, v))
}

// GpuCostLTE applies the LTE predicate on the "gpuCost" field.
func GpuCostLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldGpuCost, v))
}

// GpuCountEQ applies the EQ predicate on the "gpuCount" field.
func GpuCountEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldGpuCount, v))
}

// GpuCountNEQ applies the NEQ predicate on the "gpuCount" field.
func GpuCountNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldGpuCount, v))
}

// GpuCountIn applies the In predicate on the "gpuCount" field.
func GpuCountIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldGpuCount, vs...))
}

// GpuCountNotIn applies the NotIn predicate on the "gpuCount" field.
func GpuCountNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldGpuCount, vs...))
}

// GpuCountGT applies the GT predicate on the "gpuCount" field.
func GpuCountGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldGpuCount, v))
}

// GpuCountGTE applies the GTE predicate on the "gpuCount" field.
func GpuCountGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldGpuCount, v))
}

// GpuCountLT applies the LT predicate on the "gpuCount" field.
func GpuCountLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldGpuCount, v))
}

// GpuCountLTE applies the LTE predicate on the "gpuCount" field.
func GpuCountLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldGpuCount, v))
}

// RamCostEQ applies the EQ predicate on the "ramCost" field.
func RamCostEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamCost, v))
}

// RamCostNEQ applies the NEQ predicate on the "ramCost" field.
func RamCostNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldRamCost, v))
}

// RamCostIn applies the In predicate on the "ramCost" field.
func RamCostIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldRamCost, vs...))
}

// RamCostNotIn applies the NotIn predicate on the "ramCost" field.
func RamCostNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldRamCost, vs...))
}

// RamCostGT applies the GT predicate on the "ramCost" field.
func RamCostGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldRamCost, v))
}

// RamCostGTE applies the GTE predicate on the "ramCost" field.
func RamCostGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldRamCost, v))
}

// RamCostLT applies the LT predicate on the "ramCost" field.
func RamCostLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldRamCost, v))
}

// RamCostLTE applies the LTE predicate on the "ramCost" field.
func RamCostLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldRamCost, v))
}

// RamByteRequestEQ applies the EQ predicate on the "ramByteRequest" field.
func RamByteRequestEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteRequest, v))
}

// RamByteRequestNEQ applies the NEQ predicate on the "ramByteRequest" field.
func RamByteRequestNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldRamByteRequest, v))
}

// RamByteRequestIn applies the In predicate on the "ramByteRequest" field.
func RamByteRequestIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldRamByteRequest, vs...))
}

// RamByteRequestNotIn applies the NotIn predicate on the "ramByteRequest" field.
func RamByteRequestNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldRamByteRequest, vs...))
}

// RamByteRequestGT applies the GT predicate on the "ramByteRequest" field.
func RamByteRequestGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldRamByteRequest, v))
}

// RamByteRequestGTE applies the GTE predicate on the "ramByteRequest" field.
func RamByteRequestGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldRamByteRequest, v))
}

// RamByteRequestLT applies the LT predicate on the "ramByteRequest" field.
func RamByteRequestLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldRamByteRequest, v))
}

// RamByteRequestLTE applies the LTE predicate on the "ramByteRequest" field.
func RamByteRequestLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldRamByteRequest, v))
}

// PvCostEQ applies the EQ predicate on the "pvCost" field.
func PvCostEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPvCost, v))
}

// PvCostNEQ applies the NEQ predicate on the "pvCost" field.
func PvCostNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldPvCost, v))
}

// PvCostIn applies the In predicate on the "pvCost" field.
func PvCostIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldPvCost, vs...))
}

// PvCostNotIn applies the NotIn predicate on the "pvCost" field.
func PvCostNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldPvCost, vs...))
}

// PvCostGT applies the GT predicate on the "pvCost" field.
func PvCostGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldPvCost, v))
}

// PvCostGTE applies the GTE predicate on the "pvCost" field.
func PvCostGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldPvCost, v))
}

// PvCostLT applies the LT predicate on the "pvCost" field.
func PvCostLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldPvCost, v))
}

// PvCostLTE applies the LTE predicate on the "pvCost" field.
func PvCostLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldPvCost, v))
}

// PvBytesEQ applies the EQ predicate on the "pvBytes" field.
func PvBytesEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldPvBytes, v))
}

// PvBytesNEQ applies the NEQ predicate on the "pvBytes" field.
func PvBytesNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldPvBytes, v))
}

// PvBytesIn applies the In predicate on the "pvBytes" field.
func PvBytesIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldPvBytes, vs...))
}

// PvBytesNotIn applies the NotIn predicate on the "pvBytes" field.
func PvBytesNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldPvBytes, vs...))
}

// PvBytesGT applies the GT predicate on the "pvBytes" field.
func PvBytesGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldPvBytes, v))
}

// PvBytesGTE applies the GTE predicate on the "pvBytes" field.
func PvBytesGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldPvBytes, v))
}

// PvBytesLT applies the LT predicate on the "pvBytes" field.
func PvBytesLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldPvBytes, v))
}

// PvBytesLTE applies the LTE predicate on the "pvBytes" field.
func PvBytesLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldPvBytes, v))
}

// CpuCoreUsageAverageEQ applies the EQ predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageAverageNEQ applies the NEQ predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageAverageIn applies the In predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldCpuCoreUsageAverage, vs...))
}

// CpuCoreUsageAverageNotIn applies the NotIn predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldCpuCoreUsageAverage, vs...))
}

// CpuCoreUsageAverageGT applies the GT predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageAverageGTE applies the GTE predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageAverageLT applies the LT predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageAverageLTE applies the LTE predicate on the "cpuCoreUsageAverage" field.
func CpuCoreUsageAverageLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldCpuCoreUsageAverage, v))
}

// CpuCoreUsageMaxEQ applies the EQ predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldCpuCoreUsageMax, v))
}

// CpuCoreUsageMaxNEQ applies the NEQ predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldCpuCoreUsageMax, v))
}

// CpuCoreUsageMaxIn applies the In predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldCpuCoreUsageMax, vs...))
}

// CpuCoreUsageMaxNotIn applies the NotIn predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldCpuCoreUsageMax, vs...))
}

// CpuCoreUsageMaxGT applies the GT predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldCpuCoreUsageMax, v))
}

// CpuCoreUsageMaxGTE applies the GTE predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldCpuCoreUsageMax, v))
}

// CpuCoreUsageMaxLT applies the LT predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldCpuCoreUsageMax, v))
}

// CpuCoreUsageMaxLTE applies the LTE predicate on the "cpuCoreUsageMax" field.
func CpuCoreUsageMaxLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldCpuCoreUsageMax, v))
}

// RamByteUsageAverageEQ applies the EQ predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteUsageAverage, v))
}

// RamByteUsageAverageNEQ applies the NEQ predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldRamByteUsageAverage, v))
}

// RamByteUsageAverageIn applies the In predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldRamByteUsageAverage, vs...))
}

// RamByteUsageAverageNotIn applies the NotIn predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldRamByteUsageAverage, vs...))
}

// RamByteUsageAverageGT applies the GT predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldRamByteUsageAverage, v))
}

// RamByteUsageAverageGTE applies the GTE predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldRamByteUsageAverage, v))
}

// RamByteUsageAverageLT applies the LT predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldRamByteUsageAverage, v))
}

// RamByteUsageAverageLTE applies the LTE predicate on the "ramByteUsageAverage" field.
func RamByteUsageAverageLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldRamByteUsageAverage, v))
}

// RamByteUsageMaxEQ applies the EQ predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldEQ(FieldRamByteUsageMax, v))
}

// RamByteUsageMaxNEQ applies the NEQ predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxNEQ(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNEQ(FieldRamByteUsageMax, v))
}

// RamByteUsageMaxIn applies the In predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldIn(FieldRamByteUsageMax, vs...))
}

// RamByteUsageMaxNotIn applies the NotIn predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxNotIn(vs ...float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldNotIn(FieldRamByteUsageMax, vs...))
}

// RamByteUsageMaxGT applies the GT predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxGT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGT(FieldRamByteUsageMax, v))
}

// RamByteUsageMaxGTE applies the GTE predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxGTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldGTE(FieldRamByteUsageMax, v))
}

// RamByteUsageMaxLT applies the LT predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxLT(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLT(FieldRamByteUsageMax, v))
}

// RamByteUsageMaxLTE applies the LTE predicate on the "ramByteUsageMax" field.
func RamByteUsageMaxLTE(v float64) predicate.AllocationCost {
	return predicate.AllocationCost(sql.FieldLTE(FieldRamByteUsageMax, v))
}

// HasConnector applies the HasEdge predicate on the "connector" edge.
func HasConnector() predicate.AllocationCost {
	return predicate.AllocationCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Connector
		step.Edge.Schema = schemaConfig.AllocationCost
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasConnectorWith applies the HasEdge predicate on the "connector" edge with a given conditions (other predicates).
func HasConnectorWith(preds ...predicate.Connector) predicate.AllocationCost {
	return predicate.AllocationCost(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ConnectorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ConnectorTable, ConnectorColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Connector
		step.Edge.Schema = schemaConfig.AllocationCost
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AllocationCost) predicate.AllocationCost {
	return predicate.AllocationCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AllocationCost) predicate.AllocationCost {
	return predicate.AllocationCost(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AllocationCost) predicate.AllocationCost {
	return predicate.AllocationCost(func(s *sql.Selector) {
		p(s.Not())
	})
}
