package sourcecontrolsyncjobstreams

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StreamType string

const (
	StreamTypeError  StreamType = "Error"
	StreamTypeOutput StreamType = "Output"
)

func PossibleValuesForStreamType() []string {
	return []string{
		string(StreamTypeError),
		string(StreamTypeOutput),
	}
}

func parseStreamType(input string) (*StreamType, error) {
	vals := map[string]StreamType{
		"error":  StreamTypeError,
		"output": StreamTypeOutput,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StreamType(input)
	return &out, nil
}
