package types

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgRevealSolution{}

// RevealSolutionType const for type
const RevealSolutionType = "RevealSolution"

// MsgRevealSolution - struct for unjailing jailed validator
type MsgRevealSolution struct {
	Scavenger    sdk.AccAddress `json:"scavenger" yaml:"scavenger"`       // address of the scavenger scavenger
	SolutionHash string         `json:"solutionHash" yaml:"solutionHash"` // SolutionHash of the scavenge
	Solution     string         `json:"solution" yaml:"solution"`         // solution of the scavenge
}

// NewMsgRevealSolution constructor
func NewMsgRevealSolution(scavenger sdk.AccAddress, solution string) MsgRevealSolution {
	// important
	solutionHash := sha256.Sum256([]byte(solution))
	solutionHashString := hex.EncodeToString(solutionHash[:])
	return MsgRevealSolution{
		Scavenger:    scavenger,
		SolutionHash: solutionHashString,
		Solution:     solution,
	}
}

// Route Route impl
func (msg MsgRevealSolution) Route() string { return RouterKey }

// Type Type impl
func (msg MsgRevealSolution) Type() string { return RevealSolutionType }

// GetSigners GetSigners impl
func (msg MsgRevealSolution) GetSigners() []sdk.AccAddress { return []sdk.AccAddress{msg.Scavenger} }

// GetSignBytes GetSignBytes impl
func (msg MsgRevealSolution) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(msg))
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgRevealSolution) ValidateBasic() error {
	if msg.Scavenger.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "create is empty")
	}

	if msg.Solution == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solution is empty")
	}

	if msg.SolutionHash == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "solution hash is empty")
	}

	solutionHash := sha256.Sum256([]byte(msg.Solution))
	solutionHashString := hex.EncodeToString(solutionHash[:])
	if msg.SolutionHash != solutionHashString {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("Hash of solution (%s) doesn't equal solutionHash (%s)", msg.SolutionHash, solutionHashString))
	}

	return nil
}
