package triggerruns

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RunQueryFilterOperand string

const (
	RunQueryFilterOperandActivityName        RunQueryFilterOperand = "ActivityName"
	RunQueryFilterOperandActivityRunEnd      RunQueryFilterOperand = "ActivityRunEnd"
	RunQueryFilterOperandActivityRunStart    RunQueryFilterOperand = "ActivityRunStart"
	RunQueryFilterOperandActivityType        RunQueryFilterOperand = "ActivityType"
	RunQueryFilterOperandLatestOnly          RunQueryFilterOperand = "LatestOnly"
	RunQueryFilterOperandPipelineName        RunQueryFilterOperand = "PipelineName"
	RunQueryFilterOperandRunEnd              RunQueryFilterOperand = "RunEnd"
	RunQueryFilterOperandRunGroupId          RunQueryFilterOperand = "RunGroupId"
	RunQueryFilterOperandRunStart            RunQueryFilterOperand = "RunStart"
	RunQueryFilterOperandStatus              RunQueryFilterOperand = "Status"
	RunQueryFilterOperandTriggerName         RunQueryFilterOperand = "TriggerName"
	RunQueryFilterOperandTriggerRunTimestamp RunQueryFilterOperand = "TriggerRunTimestamp"
)

func PossibleValuesForRunQueryFilterOperand() []string {
	return []string{
		string(RunQueryFilterOperandActivityName),
		string(RunQueryFilterOperandActivityRunEnd),
		string(RunQueryFilterOperandActivityRunStart),
		string(RunQueryFilterOperandActivityType),
		string(RunQueryFilterOperandLatestOnly),
		string(RunQueryFilterOperandPipelineName),
		string(RunQueryFilterOperandRunEnd),
		string(RunQueryFilterOperandRunGroupId),
		string(RunQueryFilterOperandRunStart),
		string(RunQueryFilterOperandStatus),
		string(RunQueryFilterOperandTriggerName),
		string(RunQueryFilterOperandTriggerRunTimestamp),
	}
}

func parseRunQueryFilterOperand(input string) (*RunQueryFilterOperand, error) {
	vals := map[string]RunQueryFilterOperand{
		"activityname":        RunQueryFilterOperandActivityName,
		"activityrunend":      RunQueryFilterOperandActivityRunEnd,
		"activityrunstart":    RunQueryFilterOperandActivityRunStart,
		"activitytype":        RunQueryFilterOperandActivityType,
		"latestonly":          RunQueryFilterOperandLatestOnly,
		"pipelinename":        RunQueryFilterOperandPipelineName,
		"runend":              RunQueryFilterOperandRunEnd,
		"rungroupid":          RunQueryFilterOperandRunGroupId,
		"runstart":            RunQueryFilterOperandRunStart,
		"status":              RunQueryFilterOperandStatus,
		"triggername":         RunQueryFilterOperandTriggerName,
		"triggerruntimestamp": RunQueryFilterOperandTriggerRunTimestamp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunQueryFilterOperand(input)
	return &out, nil
}

type RunQueryFilterOperator string

const (
	RunQueryFilterOperatorEquals    RunQueryFilterOperator = "Equals"
	RunQueryFilterOperatorIn        RunQueryFilterOperator = "In"
	RunQueryFilterOperatorNotEquals RunQueryFilterOperator = "NotEquals"
	RunQueryFilterOperatorNotIn     RunQueryFilterOperator = "NotIn"
)

func PossibleValuesForRunQueryFilterOperator() []string {
	return []string{
		string(RunQueryFilterOperatorEquals),
		string(RunQueryFilterOperatorIn),
		string(RunQueryFilterOperatorNotEquals),
		string(RunQueryFilterOperatorNotIn),
	}
}

func parseRunQueryFilterOperator(input string) (*RunQueryFilterOperator, error) {
	vals := map[string]RunQueryFilterOperator{
		"equals":    RunQueryFilterOperatorEquals,
		"in":        RunQueryFilterOperatorIn,
		"notequals": RunQueryFilterOperatorNotEquals,
		"notin":     RunQueryFilterOperatorNotIn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunQueryFilterOperator(input)
	return &out, nil
}

type RunQueryOrder string

const (
	RunQueryOrderASC  RunQueryOrder = "ASC"
	RunQueryOrderDESC RunQueryOrder = "DESC"
)

func PossibleValuesForRunQueryOrder() []string {
	return []string{
		string(RunQueryOrderASC),
		string(RunQueryOrderDESC),
	}
}

func parseRunQueryOrder(input string) (*RunQueryOrder, error) {
	vals := map[string]RunQueryOrder{
		"asc":  RunQueryOrderASC,
		"desc": RunQueryOrderDESC,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunQueryOrder(input)
	return &out, nil
}

type RunQueryOrderByField string

const (
	RunQueryOrderByFieldActivityName        RunQueryOrderByField = "ActivityName"
	RunQueryOrderByFieldActivityRunEnd      RunQueryOrderByField = "ActivityRunEnd"
	RunQueryOrderByFieldActivityRunStart    RunQueryOrderByField = "ActivityRunStart"
	RunQueryOrderByFieldPipelineName        RunQueryOrderByField = "PipelineName"
	RunQueryOrderByFieldRunEnd              RunQueryOrderByField = "RunEnd"
	RunQueryOrderByFieldRunStart            RunQueryOrderByField = "RunStart"
	RunQueryOrderByFieldStatus              RunQueryOrderByField = "Status"
	RunQueryOrderByFieldTriggerName         RunQueryOrderByField = "TriggerName"
	RunQueryOrderByFieldTriggerRunTimestamp RunQueryOrderByField = "TriggerRunTimestamp"
)

func PossibleValuesForRunQueryOrderByField() []string {
	return []string{
		string(RunQueryOrderByFieldActivityName),
		string(RunQueryOrderByFieldActivityRunEnd),
		string(RunQueryOrderByFieldActivityRunStart),
		string(RunQueryOrderByFieldPipelineName),
		string(RunQueryOrderByFieldRunEnd),
		string(RunQueryOrderByFieldRunStart),
		string(RunQueryOrderByFieldStatus),
		string(RunQueryOrderByFieldTriggerName),
		string(RunQueryOrderByFieldTriggerRunTimestamp),
	}
}

func parseRunQueryOrderByField(input string) (*RunQueryOrderByField, error) {
	vals := map[string]RunQueryOrderByField{
		"activityname":        RunQueryOrderByFieldActivityName,
		"activityrunend":      RunQueryOrderByFieldActivityRunEnd,
		"activityrunstart":    RunQueryOrderByFieldActivityRunStart,
		"pipelinename":        RunQueryOrderByFieldPipelineName,
		"runend":              RunQueryOrderByFieldRunEnd,
		"runstart":            RunQueryOrderByFieldRunStart,
		"status":              RunQueryOrderByFieldStatus,
		"triggername":         RunQueryOrderByFieldTriggerName,
		"triggerruntimestamp": RunQueryOrderByFieldTriggerRunTimestamp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RunQueryOrderByField(input)
	return &out, nil
}

type TriggerRunStatus string

const (
	TriggerRunStatusFailed     TriggerRunStatus = "Failed"
	TriggerRunStatusInprogress TriggerRunStatus = "Inprogress"
	TriggerRunStatusSucceeded  TriggerRunStatus = "Succeeded"
)

func PossibleValuesForTriggerRunStatus() []string {
	return []string{
		string(TriggerRunStatusFailed),
		string(TriggerRunStatusInprogress),
		string(TriggerRunStatusSucceeded),
	}
}

func parseTriggerRunStatus(input string) (*TriggerRunStatus, error) {
	vals := map[string]TriggerRunStatus{
		"failed":     TriggerRunStatusFailed,
		"inprogress": TriggerRunStatusInprogress,
		"succeeded":  TriggerRunStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerRunStatus(input)
	return &out, nil
}
