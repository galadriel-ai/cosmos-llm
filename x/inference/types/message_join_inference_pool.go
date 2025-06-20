package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgJoinInferencePool{}

func NewMsgJoinInferencePool(creator string, modelId uint64) *MsgJoinInferencePool {
	return &MsgJoinInferencePool{
		Creator: creator,
		ModelId: modelId,
	}
}

func (msg *MsgJoinInferencePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
