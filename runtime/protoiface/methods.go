// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package protoiface contains types referenced or implemented by messages.
//
// WARNING: This package should only be imported by message implementations.
// The functionality found in this package should be accessed through
// higher-level abstractions provided by the proto package.
package protoiface

import (
	"google.golang.org/protobuf/internal/pragma"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// Methods is a set of optional fast-path implementations of various operations.
type Methods = struct {
	pragma.NoUnkeyedLiterals

	// Flags indicate support for optional features.
	Flags SupportFlags

	// Size returns the size in bytes of the wire-format encoding of m.
	// Marshal must be provided if a custom Size is provided.
	Size func(SizeInput) SizeOutput

	// Marshal writes the wire-format encoding of m to the provided buffer.
	// Size should be provided if a custom Marshal is provided.
	// It must not return an error for a partial message.
	Marshal func(MarshalInput) (MarshalOutput, error)

	// Unmarshal parses the wire-format encoding of a message and merges the result to m.
	// It must not reset the target message or return an error for a partial message.
	Unmarshal func(UnmarshalInput) (UnmarshalOutput, error)

	// IsInitialized returns an error if any required fields in m are not set.
	IsInitialized func(IsInitializedInput) (IsInitializedOutput, error)

	// Merge merges src into dst.
	Merge func(MergeInput) MergeOutput
}

// SupportFlags indicate support for optional features.
type SupportFlags = uint64

const (
	// SupportMarshalDeterministic reports whether MarshalOptions.Deterministic is supported.
	SupportMarshalDeterministic SupportFlags = 1 << iota

	// SupportUnmarshalDiscardUnknown reports whether UnmarshalOptions.DiscardUnknown is supported.
	SupportUnmarshalDiscardUnknown
)

// SizeInput is input to the Size method.
type SizeInput = struct {
	pragma.NoUnkeyedLiterals

	Message protoreflect.Message
	Flags   MarshalInputFlags
}

// SizeOutput is output from the Size method.
type SizeOutput = struct {
	pragma.NoUnkeyedLiterals

	Size int
}

// MarshalInput is input to the Marshal method.
type MarshalInput = struct {
	pragma.NoUnkeyedLiterals

	Message protoreflect.Message
	Buf     []byte // output is appended to this buffer
	Flags   MarshalInputFlags
}

// MarshalOutput is output from the Marshal method.
type MarshalOutput = struct {
	pragma.NoUnkeyedLiterals

	Buf []byte // contains marshaled message
}

// MarshalInputFlags configure the marshaler.
// Most flags correspond to fields in proto.MarshalOptions.
type MarshalInputFlags = uint8

const (
	MarshalDeterministic MarshalInputFlags = 1 << iota
	MarshalUseCachedSize
)

// UnmarshalInput is input to the Unmarshal method.
type UnmarshalInput = struct {
	pragma.NoUnkeyedLiterals

	Message  protoreflect.Message
	Buf      []byte // input buffer
	Flags    UnmarshalInputFlags
	Resolver interface {
		FindExtensionByName(field protoreflect.FullName) (protoreflect.ExtensionType, error)
		FindExtensionByNumber(message protoreflect.FullName, field protoreflect.FieldNumber) (protoreflect.ExtensionType, error)
	}
}

// UnmarshalOutput is output from the Unmarshal method.
type UnmarshalOutput = struct {
	pragma.NoUnkeyedLiterals

	Flags UnmarshalOutputFlags
}

// UnmarshalInputFlags configure the unmarshaler.
// Most flags correspond to fields in proto.UnmarshalOptions.
type UnmarshalInputFlags = uint8

const (
	UnmarshalDiscardUnknown UnmarshalInputFlags = 1 << iota
)

// UnmarshalOutputFlags are output from the Unmarshal method.
type UnmarshalOutputFlags = uint8

const (
	// UnmarshalInitialized may be set on return if all required fields are known to be set.
	// A value of false does not indicate that the message is uninitialized, only
	// that its status could not be confirmed.
	UnmarshalInitialized UnmarshalOutputFlags = 1 << iota
)

// IsInitializedInput is input to the IsInitialized method.
type IsInitializedInput = struct {
	pragma.NoUnkeyedLiterals

	Message protoreflect.Message
}

// IsInitializedOutput is output from the IsInitialized method.
type IsInitializedOutput = struct {
	pragma.NoUnkeyedLiterals

	Flags IsInitializedOutputFlags
}

// IsInitializedOutputFlags are output from the IsInitialized method.
type IsInitializedOutputFlags = uint8

const (
	// IsInitialized reports whether the message is initialized.
	IsInitialized IsInitializedOutputFlags = 1 << iota
)

// MergeInput is input to the Merge method.
type MergeInput = struct {
	pragma.NoUnkeyedLiterals

	Source      protoreflect.Message
	Destination protoreflect.Message
}

// MergeOutput is output from the Merge method.
type MergeOutput = struct {
	pragma.NoUnkeyedLiterals

	Flags MergeOutputFlags
}

// MergeOutputFlags are output from the Merge method.
type MergeOutputFlags = uint8

const (
	// MergeComplete reports whether the merge was performed.
	// If unset, the merger must have made no changes to the destination.
	MergeComplete MergeOutputFlags = 1 << iota
)
