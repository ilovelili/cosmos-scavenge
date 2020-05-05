package type

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCommitSolution{}

// CommitSolutionType const for type
const CommitSolutionType = "CommitSolution"

// MsgCommitSolution - struct for unjailing jailed validator
type MsgCommitSolution struct {
	Scavenger             sdk.AccAddress `json:"scavenger" yaml:"scavenger"`                         // address of the scavenger
	SolutionHash          string         `json:"solutionhash" yaml:"solutionhash"`                   // solutionhash of the scavenge
	SolutionScavengerHash string         `json:"solutionScavengerHash" yaml:"solutionScavengerHash"` // solution hash of the scavenge
}

// NewMsgCommitSolution creates a new MsgCommitSolution instance
func NewMsgCommitSolution(scavenger sdk.AccAddress, solutionHash string, solutionScavengerHash string) MsgCommitSolution {
	return MsgCommitSolution{
		Scavenger:             scavenger,
		SolutionHash:          solutionHash,
		SolutionScavengerHash: solutionScavengerHash,
	}
}

// Route Route impl
func (msg MsgCommitSolution) Route() string { return RouterKey }

// Type Type impl
func (msg MsgCommitSolution) Type() string { return CommitSolutionType }

// GetSigners GetSigners impl
func (msg MsgCommitSolution) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.Scavenger} }

// GetSignBytes GetSignBytes impl
func (msg MsgCommitSolution) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgCommitSolution) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "creator is empty")
	}

	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solution hash is empty")
	}

	if msg.SolutionScavengerHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solution scavenge hash is empty")
	}

	return nil
}
