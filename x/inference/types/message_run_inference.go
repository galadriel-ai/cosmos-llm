package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRunInference{}

func NewMsgRunInference(creator string, prompt string, modelid uint64) *MsgRunInference {
	return &MsgRunInference{
		Creator: creator,
		Prompt:  prompt,
		Modelid: modelid,
	}
}

func (msg *MsgRunInference) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
