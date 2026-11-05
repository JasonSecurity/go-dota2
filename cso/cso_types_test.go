package cso

import (
	"reflect"
	"testing"

	protocol "github.com/paralin/go-dota2/protocol"
)

func TestNewSharedObjectRegistersSupportedTypes(t *testing.T) {
	tests := []struct {
		name string
		typ  CSOType
		want interface{}
	}{
		{name: "econ item", typ: EconItem, want: &protocol.CSOEconItem{}},
		{name: "econ game account client", typ: EconGameAccountClient, want: &bgcm.CSOEconGameAccountClient{}},
		{name: "Dota game account client", typ: GameAccountClient, want: &protocol.CSODOTAGameAccountClient{}},
		{name: "party", typ: Party, want: &protocol.CSODOTAParty{}},
		{name: "lobby", typ: Lobby, want: &gcmm.CSODOTALobby{}},
		{name: "item tournament passport", typ: ItemTournamentPassport, want: &bgcm.CSOEconItemTournamentPassport{}},
		{name: "game account plus", typ: GameAccountPlus, want: &gccm.CSODOTAGameAccountPlus{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSharedObject(tt.typ)
			if err != nil {
				t.Fatalf("NewSharedObject(%d) returned error: %v", tt.typ, err)
			}
			if reflect.TypeOf(got) != reflect.TypeOf(tt.want) {
				t.Fatalf("NewSharedObject(%d) type = %T, want %T", tt.typ, got, tt.want)
			}
		})
	}
}

func TestNewSharedObjectRejectsUnknownType(t *testing.T) {
	const unknownType CSOType = 999999
	if got, err := NewSharedObject(unknownType); err == nil {
		t.Fatalf("NewSharedObject(%d) = %T, nil error; want unknown-type error", unknownType, got)
	}
}
