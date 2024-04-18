package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/inference module sentinel errors
var (
	ErrInvalidSigner     = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample            = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidModel      = sdkerrors.Register(ModuleName, 1102, "model id invalid")
	ErrInvalidAddress    = sdkerrors.Register(ModuleName, 1103, "msg creator address invalid: %s")
	ErrFailedToPay       = sdkerrors.Register(ModuleName, 1104, "failed to make payment with account address: %s")
	ErrAlreadyRegistered = sdkerrors.Register(ModuleName, 1105, "model id already registered with given account")
	ErrAlreadyResponded  = sdkerrors.Register(ModuleName, 1106, "run id already responded to with given account")
	ErrRunNotFound       = sdkerrors.Register(ModuleName, 1107, "run id for inference responses not found")
	ErrNoUnansweredRuns  = sdkerrors.Register(ModuleName, 1108, "model id has no unanswered inference runs")
)
