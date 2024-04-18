package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/inference module sentinel errors
var (
	ErrInvalidSigner  = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrSample         = sdkerrors.Register(ModuleName, 1101, "sample error")
	ErrInvalidModel   = sdkerrors.Register(ModuleName, 1102, "model id invalid")
	ErrInvalidAddress = sdkerrors.Register(ModuleName, 1103, "msg creator address invalid: %s")
	ErrFailedToPay    = sdkerrors.Register(ModuleName, 1104, "failed to make payment with account address: %s")
)
