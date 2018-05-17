// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto

/*
Package dpos is a generated protocol buffer package.

It is generated from these files:
	github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto

It has these top-level messages:
	Params
	State
	Voter
	Witness
	Candidate
	CandidateSet
	Vote
	VoteSet
	InitRequest
	RegisterCandidateRequest
	RegisterCandidateResponse
	UnregisterCandidateRequest
	UnregisterCandidateResponse
	VoteRequest
	VoteResponse
	ProxyVoteRequest
	ProxyVoteResponse
	UnproxyVoteRequest
	UnproxyVoteResponse
	ElectRequest
	ElectResponse
	ListWitnessesRequest
	ListWitnessesResponse
*/
package dpos

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/loomnetwork/go-loom/types"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Params struct {
	WitnessCount        uint64         `protobuf:"varint,1,opt,name=witness_count,json=witnessCount,proto3" json:"witness_count,omitempty"`
	VoteAllocation      uint64         `protobuf:"varint,2,opt,name=vote_allocation,json=voteAllocation,proto3" json:"vote_allocation,omitempty"`
	TermLength          uint64         `protobuf:"varint,3,opt,name=term_length,json=termLength,proto3" json:"term_length,omitempty"`
	MinPowerFraction    uint64         `protobuf:"varint,4,opt,name=min_power_fraction,json=minPowerFraction,proto3" json:"min_power_fraction,omitempty"`
	CoinContractAddress *types.Address `protobuf:"bytes,5,opt,name=coin_contract_address,json=coinContractAddress" json:"coin_contract_address,omitempty"`
}

func (m *Params) Reset()                    { *m = Params{} }
func (m *Params) String() string            { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()               {}
func (*Params) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{0} }

func (m *Params) GetWitnessCount() uint64 {
	if m != nil {
		return m.WitnessCount
	}
	return 0
}

func (m *Params) GetVoteAllocation() uint64 {
	if m != nil {
		return m.VoteAllocation
	}
	return 0
}

func (m *Params) GetTermLength() uint64 {
	if m != nil {
		return m.TermLength
	}
	return 0
}

func (m *Params) GetMinPowerFraction() uint64 {
	if m != nil {
		return m.MinPowerFraction
	}
	return 0
}

func (m *Params) GetCoinContractAddress() *types.Address {
	if m != nil {
		return m.CoinContractAddress
	}
	return nil
}

type State struct {
	Params       *Params    `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	LastElection uint64     `protobuf:"varint,2,opt,name=last_election,json=lastElection,proto3" json:"last_election,omitempty"`
	Witnesses    []*Witness `protobuf:"bytes,3,rep,name=witnesses" json:"witnesses,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{1} }

func (m *State) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *State) GetLastElection() uint64 {
	if m != nil {
		return m.LastElection
	}
	return 0
}

func (m *State) GetWitnesses() []*Witness {
	if m != nil {
		return m.Witnesses
	}
	return nil
}

type Voter struct {
	Address            *types.Address   `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Balance            uint64           `protobuf:"varint,2,opt,name=balance,proto3" json:"balance,omitempty"`
	ProxyAddress       *types.Address   `protobuf:"bytes,3,opt,name=proxy_address,json=proxyAddress" json:"proxy_address,omitempty"`
	PrincipalAddresses []*types.Address `protobuf:"bytes,4,rep,name=principal_addresses,json=principalAddresses" json:"principal_addresses,omitempty"`
}

func (m *Voter) Reset()                    { *m = Voter{} }
func (m *Voter) String() string            { return proto.CompactTextString(m) }
func (*Voter) ProtoMessage()               {}
func (*Voter) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{2} }

func (m *Voter) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Voter) GetBalance() uint64 {
	if m != nil {
		return m.Balance
	}
	return 0
}

func (m *Voter) GetProxyAddress() *types.Address {
	if m != nil {
		return m.ProxyAddress
	}
	return nil
}

func (m *Voter) GetPrincipalAddresses() []*types.Address {
	if m != nil {
		return m.PrincipalAddresses
	}
	return nil
}

type Witness struct {
	PubKey     []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	VoteTotal  uint64 `protobuf:"varint,2,opt,name=vote_total,json=voteTotal,proto3" json:"vote_total,omitempty"`
	PowerTotal uint64 `protobuf:"varint,3,opt,name=power_total,json=powerTotal,proto3" json:"power_total,omitempty"`
}

func (m *Witness) Reset()                    { *m = Witness{} }
func (m *Witness) String() string            { return proto.CompactTextString(m) }
func (*Witness) ProtoMessage()               {}
func (*Witness) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{3} }

func (m *Witness) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

func (m *Witness) GetVoteTotal() uint64 {
	if m != nil {
		return m.VoteTotal
	}
	return 0
}

func (m *Witness) GetPowerTotal() uint64 {
	if m != nil {
		return m.PowerTotal
	}
	return 0
}

type Candidate struct {
	Address *types.Address `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	PubKey  []byte         `protobuf:"bytes,2,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *Candidate) Reset()                    { *m = Candidate{} }
func (m *Candidate) String() string            { return proto.CompactTextString(m) }
func (*Candidate) ProtoMessage()               {}
func (*Candidate) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{4} }

func (m *Candidate) GetAddress() *types.Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Candidate) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type CandidateSet struct {
	Candidates map[string]*Candidate `protobuf:"bytes,1,rep,name=candidates" json:"candidates,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *CandidateSet) Reset()                    { *m = CandidateSet{} }
func (m *CandidateSet) String() string            { return proto.CompactTextString(m) }
func (*CandidateSet) ProtoMessage()               {}
func (*CandidateSet) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{5} }

func (m *CandidateSet) GetCandidates() map[string]*Candidate {
	if m != nil {
		return m.Candidates
	}
	return nil
}

type Vote struct {
	VoterAddress     *types.Address `protobuf:"bytes,1,opt,name=voter_address,json=voterAddress" json:"voter_address,omitempty"`
	CandidateAddress *types.Address `protobuf:"bytes,2,opt,name=candidate_address,json=candidateAddress" json:"candidate_address,omitempty"`
	Amount           uint64         `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Vote) Reset()                    { *m = Vote{} }
func (m *Vote) String() string            { return proto.CompactTextString(m) }
func (*Vote) ProtoMessage()               {}
func (*Vote) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{6} }

func (m *Vote) GetVoterAddress() *types.Address {
	if m != nil {
		return m.VoterAddress
	}
	return nil
}

func (m *Vote) GetCandidateAddress() *types.Address {
	if m != nil {
		return m.CandidateAddress
	}
	return nil
}

func (m *Vote) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteSet struct {
	Votes map[string]*Vote `protobuf:"bytes,1,rep,name=votes" json:"votes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *VoteSet) Reset()                    { *m = VoteSet{} }
func (m *VoteSet) String() string            { return proto.CompactTextString(m) }
func (*VoteSet) ProtoMessage()               {}
func (*VoteSet) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{7} }

func (m *VoteSet) GetVotes() map[string]*Vote {
	if m != nil {
		return m.Votes
	}
	return nil
}

type InitRequest struct {
	Params     *Params            `protobuf:"bytes,1,opt,name=params" json:"params,omitempty"`
	Validators []*types.Validator `protobuf:"bytes,2,rep,name=validators" json:"validators,omitempty"`
}

func (m *InitRequest) Reset()                    { *m = InitRequest{} }
func (m *InitRequest) String() string            { return proto.CompactTextString(m) }
func (*InitRequest) ProtoMessage()               {}
func (*InitRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{8} }

func (m *InitRequest) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *InitRequest) GetValidators() []*types.Validator {
	if m != nil {
		return m.Validators
	}
	return nil
}

type RegisterCandidateRequest struct {
	PubKey []byte `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
}

func (m *RegisterCandidateRequest) Reset()                    { *m = RegisterCandidateRequest{} }
func (m *RegisterCandidateRequest) String() string            { return proto.CompactTextString(m) }
func (*RegisterCandidateRequest) ProtoMessage()               {}
func (*RegisterCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{9} }

func (m *RegisterCandidateRequest) GetPubKey() []byte {
	if m != nil {
		return m.PubKey
	}
	return nil
}

type RegisterCandidateResponse struct {
}

func (m *RegisterCandidateResponse) Reset()                    { *m = RegisterCandidateResponse{} }
func (m *RegisterCandidateResponse) String() string            { return proto.CompactTextString(m) }
func (*RegisterCandidateResponse) ProtoMessage()               {}
func (*RegisterCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{10} }

type UnregisterCandidateRequest struct {
}

func (m *UnregisterCandidateRequest) Reset()                    { *m = UnregisterCandidateRequest{} }
func (m *UnregisterCandidateRequest) String() string            { return proto.CompactTextString(m) }
func (*UnregisterCandidateRequest) ProtoMessage()               {}
func (*UnregisterCandidateRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{11} }

type UnregisterCandidateResponse struct {
}

func (m *UnregisterCandidateResponse) Reset()                    { *m = UnregisterCandidateResponse{} }
func (m *UnregisterCandidateResponse) String() string            { return proto.CompactTextString(m) }
func (*UnregisterCandidateResponse) ProtoMessage()               {}
func (*UnregisterCandidateResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{12} }

type VoteRequest struct {
	CandidateAddress *types.Address `protobuf:"bytes,1,opt,name=candidate_address,json=candidateAddress" json:"candidate_address,omitempty"`
	Amount           int64          `protobuf:"varint,2,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *VoteRequest) Reset()                    { *m = VoteRequest{} }
func (m *VoteRequest) String() string            { return proto.CompactTextString(m) }
func (*VoteRequest) ProtoMessage()               {}
func (*VoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{13} }

func (m *VoteRequest) GetCandidateAddress() *types.Address {
	if m != nil {
		return m.CandidateAddress
	}
	return nil
}

func (m *VoteRequest) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type VoteResponse struct {
}

func (m *VoteResponse) Reset()                    { *m = VoteResponse{} }
func (m *VoteResponse) String() string            { return proto.CompactTextString(m) }
func (*VoteResponse) ProtoMessage()               {}
func (*VoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{14} }

type ProxyVoteRequest struct {
	ProxyAddress *types.Address `protobuf:"bytes,1,opt,name=proxy_address,json=proxyAddress" json:"proxy_address,omitempty"`
}

func (m *ProxyVoteRequest) Reset()                    { *m = ProxyVoteRequest{} }
func (m *ProxyVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*ProxyVoteRequest) ProtoMessage()               {}
func (*ProxyVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{15} }

func (m *ProxyVoteRequest) GetProxyAddress() *types.Address {
	if m != nil {
		return m.ProxyAddress
	}
	return nil
}

type ProxyVoteResponse struct {
}

func (m *ProxyVoteResponse) Reset()                    { *m = ProxyVoteResponse{} }
func (m *ProxyVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*ProxyVoteResponse) ProtoMessage()               {}
func (*ProxyVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{16} }

type UnproxyVoteRequest struct {
}

func (m *UnproxyVoteRequest) Reset()                    { *m = UnproxyVoteRequest{} }
func (m *UnproxyVoteRequest) String() string            { return proto.CompactTextString(m) }
func (*UnproxyVoteRequest) ProtoMessage()               {}
func (*UnproxyVoteRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{17} }

type UnproxyVoteResponse struct {
}

func (m *UnproxyVoteResponse) Reset()                    { *m = UnproxyVoteResponse{} }
func (m *UnproxyVoteResponse) String() string            { return proto.CompactTextString(m) }
func (*UnproxyVoteResponse) ProtoMessage()               {}
func (*UnproxyVoteResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{18} }

type ElectRequest struct {
}

func (m *ElectRequest) Reset()                    { *m = ElectRequest{} }
func (m *ElectRequest) String() string            { return proto.CompactTextString(m) }
func (*ElectRequest) ProtoMessage()               {}
func (*ElectRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{19} }

type ElectResponse struct {
}

func (m *ElectResponse) Reset()                    { *m = ElectResponse{} }
func (m *ElectResponse) String() string            { return proto.CompactTextString(m) }
func (*ElectResponse) ProtoMessage()               {}
func (*ElectResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{20} }

type ListWitnessesRequest struct {
}

func (m *ListWitnessesRequest) Reset()                    { *m = ListWitnessesRequest{} }
func (m *ListWitnessesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListWitnessesRequest) ProtoMessage()               {}
func (*ListWitnessesRequest) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{21} }

type ListWitnessesResponse struct {
	Witnesses []*Witness `protobuf:"bytes,1,rep,name=witnesses" json:"witnesses,omitempty"`
}

func (m *ListWitnessesResponse) Reset()                    { *m = ListWitnessesResponse{} }
func (m *ListWitnessesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListWitnessesResponse) ProtoMessage()               {}
func (*ListWitnessesResponse) Descriptor() ([]byte, []int) { return fileDescriptorDpos, []int{22} }

func (m *ListWitnessesResponse) GetWitnesses() []*Witness {
	if m != nil {
		return m.Witnesses
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "Params")
	proto.RegisterType((*State)(nil), "State")
	proto.RegisterType((*Voter)(nil), "Voter")
	proto.RegisterType((*Witness)(nil), "Witness")
	proto.RegisterType((*Candidate)(nil), "Candidate")
	proto.RegisterType((*CandidateSet)(nil), "CandidateSet")
	proto.RegisterType((*Vote)(nil), "Vote")
	proto.RegisterType((*VoteSet)(nil), "VoteSet")
	proto.RegisterType((*InitRequest)(nil), "InitRequest")
	proto.RegisterType((*RegisterCandidateRequest)(nil), "RegisterCandidateRequest")
	proto.RegisterType((*RegisterCandidateResponse)(nil), "RegisterCandidateResponse")
	proto.RegisterType((*UnregisterCandidateRequest)(nil), "UnregisterCandidateRequest")
	proto.RegisterType((*UnregisterCandidateResponse)(nil), "UnregisterCandidateResponse")
	proto.RegisterType((*VoteRequest)(nil), "VoteRequest")
	proto.RegisterType((*VoteResponse)(nil), "VoteResponse")
	proto.RegisterType((*ProxyVoteRequest)(nil), "ProxyVoteRequest")
	proto.RegisterType((*ProxyVoteResponse)(nil), "ProxyVoteResponse")
	proto.RegisterType((*UnproxyVoteRequest)(nil), "UnproxyVoteRequest")
	proto.RegisterType((*UnproxyVoteResponse)(nil), "UnproxyVoteResponse")
	proto.RegisterType((*ElectRequest)(nil), "ElectRequest")
	proto.RegisterType((*ElectResponse)(nil), "ElectResponse")
	proto.RegisterType((*ListWitnessesRequest)(nil), "ListWitnessesRequest")
	proto.RegisterType((*ListWitnessesResponse)(nil), "ListWitnessesResponse")
}

func init() {
	proto.RegisterFile("github.com/loomnetwork/go-loom/builtin/types/dpos/dpos.proto", fileDescriptorDpos)
}

var fileDescriptorDpos = []byte{
	// 772 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x7f, 0x4f, 0xd3, 0x40,
	0x18, 0x4e, 0xf7, 0xd3, 0xbd, 0x2b, 0x30, 0x6e, 0x80, 0x75, 0x48, 0x58, 0x6a, 0xa2, 0x68, 0x64,
	0x33, 0x10, 0x13, 0x35, 0x18, 0xb2, 0x10, 0x8c, 0x44, 0xfe, 0x20, 0x45, 0x20, 0x31, 0x26, 0xcd,
	0xad, 0x3b, 0x47, 0x43, 0x77, 0x57, 0xdb, 0x2b, 0xb8, 0xff, 0xfd, 0x10, 0x7e, 0x0b, 0x3f, 0x94,
	0x5f, 0xc4, 0xdc, 0xf5, 0xae, 0x2b, 0x73, 0x13, 0xfd, 0xa7, 0xf0, 0x3e, 0xef, 0xf3, 0xbe, 0xf7,
	0x3e, 0xcf, 0xdb, 0xde, 0x60, 0x6f, 0xe8, 0xf3, 0xcb, 0xa4, 0xdf, 0xf1, 0xd8, 0xa8, 0x1b, 0x30,
	0x36, 0xa2, 0x84, 0xdf, 0xb0, 0xe8, 0xaa, 0x3b, 0x64, 0xdb, 0x22, 0xec, 0xf6, 0x13, 0x3f, 0xe0,
	0x3e, 0xed, 0xf2, 0x71, 0x48, 0xe2, 0xee, 0x20, 0x64, 0xe9, 0xa3, 0x13, 0x46, 0x8c, 0xb3, 0xd6,
	0x8b, 0x3b, 0xaa, 0xd3, 0x2a, 0xf9, 0x4c, 0x2b, 0xec, 0x5f, 0x06, 0x54, 0x4e, 0x70, 0x84, 0x47,
	0x31, 0x7a, 0x04, 0x0b, 0x37, 0x3e, 0xa7, 0x24, 0x8e, 0x5d, 0x8f, 0x25, 0x94, 0x5b, 0x46, 0xdb,
	0xd8, 0x2a, 0x39, 0xa6, 0x02, 0x0f, 0x04, 0x86, 0x9e, 0xc0, 0xd2, 0x35, 0xe3, 0xc4, 0xc5, 0x41,
	0xc0, 0x3c, 0xcc, 0x7d, 0x46, 0xad, 0x82, 0xa4, 0x2d, 0x0a, 0xb8, 0x97, 0xa1, 0x68, 0x13, 0xea,
	0x9c, 0x44, 0x23, 0x37, 0x20, 0x74, 0xc8, 0x2f, 0xad, 0xa2, 0x24, 0x81, 0x80, 0x8e, 0x25, 0x82,
	0x9e, 0x03, 0x1a, 0xf9, 0xd4, 0x0d, 0xd9, 0x0d, 0x89, 0xdc, 0x2f, 0x11, 0xf6, 0x64, 0xb3, 0x92,
	0xe4, 0x35, 0x46, 0x3e, 0x3d, 0x11, 0x89, 0x77, 0x0a, 0x47, 0x7b, 0xb0, 0xea, 0x31, 0x9f, 0xba,
	0x1e, 0xa3, 0x5c, 0x60, 0x2e, 0x1e, 0x0c, 0x22, 0x12, 0xc7, 0x56, 0xb9, 0x6d, 0x6c, 0xd5, 0x77,
	0xee, 0x75, 0x7a, 0x69, 0xec, 0x34, 0x05, 0xed, 0x40, 0xb1, 0x14, 0x68, 0x27, 0x50, 0x3e, 0xe5,
	0x98, 0x13, 0xb4, 0x09, 0x95, 0x50, 0xaa, 0x95, 0xe2, 0xea, 0x3b, 0xd5, 0x4e, 0x2a, 0xde, 0x51,
	0xb0, 0x30, 0x21, 0xc0, 0x31, 0x77, 0x49, 0x40, 0xbc, 0x9c, 0x3a, 0x53, 0x80, 0x87, 0x0a, 0x43,
	0x8f, 0xa1, 0xa6, 0x4c, 0x21, 0xb1, 0x55, 0x6c, 0x17, 0xe5, 0x00, 0x17, 0x29, 0xe2, 0x4c, 0x52,
	0xf6, 0x4f, 0x03, 0xca, 0xe7, 0x8c, 0x93, 0x08, 0xd9, 0x50, 0xd5, 0x03, 0x1b, 0x53, 0x03, 0xeb,
	0x04, 0xb2, 0xa0, 0xda, 0xc7, 0x01, 0xa6, 0x1e, 0x51, 0x87, 0xea, 0x10, 0x6d, 0xc3, 0x42, 0x18,
	0xb1, 0x6f, 0xe3, 0x4c, 0x74, 0x71, 0xaa, 0x87, 0x29, 0xd3, 0x2a, 0x42, 0xaf, 0xa1, 0x19, 0x46,
	0x3e, 0xf5, 0xfc, 0x10, 0x07, 0xba, 0x84, 0xc4, 0x56, 0x49, 0x0d, 0xaa, 0x8b, 0x50, 0x46, 0xea,
	0x69, 0x8e, 0xdd, 0x87, 0xaa, 0xd2, 0x81, 0xee, 0x43, 0x35, 0x4c, 0xfa, 0xee, 0x15, 0x19, 0xcb,
	0x91, 0x4d, 0xa7, 0x12, 0x26, 0xfd, 0x0f, 0x64, 0x8c, 0x36, 0x00, 0xe4, 0x2b, 0xc0, 0x19, 0xc7,
	0x81, 0x1a, 0xb5, 0x26, 0x90, 0x8f, 0x02, 0x10, 0x8b, 0x4f, 0x77, 0x9a, 0xe6, 0xd5, 0xe2, 0x25,
	0x24, 0x09, 0xf6, 0x7b, 0xa8, 0x1d, 0x60, 0x3a, 0xf0, 0x07, 0x62, 0x21, 0xff, 0x62, 0x4c, 0x6e,
	0x92, 0x42, 0x7e, 0x12, 0xfb, 0x87, 0x01, 0x66, 0xd6, 0xea, 0x94, 0x70, 0xf4, 0x16, 0xc0, 0xd3,
	0xb1, 0x68, 0x28, 0x04, 0x6f, 0x74, 0xf2, 0x94, 0x49, 0x10, 0x1f, 0x52, 0x1e, 0x8d, 0x9d, 0x5c,
	0x41, 0xeb, 0x08, 0x96, 0xa6, 0xd2, 0xa8, 0x01, 0x45, 0xed, 0x40, 0xcd, 0x11, 0xff, 0xa2, 0x36,
	0x94, 0xaf, 0x71, 0x90, 0xa4, 0x4b, 0xaa, 0xef, 0xc0, 0xa4, 0xa3, 0x93, 0x26, 0xde, 0x14, 0x5e,
	0x19, 0xf6, 0x77, 0x03, 0x4a, 0x62, 0xf5, 0x62, 0x77, 0xc2, 0x9b, 0xc8, 0x9d, 0x27, 0xd3, 0x94,
	0x69, 0xbd, 0xbb, 0x97, 0xb0, 0x9c, 0x0d, 0x94, 0x95, 0x14, 0xa6, 0x4a, 0x1a, 0x19, 0x45, 0x97,
	0xad, 0x41, 0x05, 0x8f, 0xe4, 0x47, 0x9b, 0xfa, 0xad, 0x22, 0x3b, 0x81, 0xaa, 0x98, 0x42, 0x78,
	0xf3, 0x14, 0xca, 0xe2, 0x24, 0x6d, 0x4b, 0xb3, 0xa3, 0x12, 0xf2, 0xaf, 0x32, 0x23, 0x65, 0xb4,
	0xf6, 0x01, 0x26, 0xe0, 0x0c, 0x0b, 0xd6, 0x6f, 0x5b, 0x50, 0x96, 0x2d, 0xf2, 0xea, 0x3f, 0x41,
	0xfd, 0x88, 0xfa, 0xdc, 0x21, 0x5f, 0x13, 0x12, 0xf3, 0xbb, 0xbf, 0xba, 0x67, 0x00, 0xd7, 0x38,
	0x10, 0x8a, 0x58, 0x24, 0xe4, 0x16, 0xa5, 0xb1, 0xe7, 0x1a, 0x72, 0x72, 0x59, 0x7b, 0x17, 0x2c,
	0x87, 0x0c, 0xfd, 0x98, 0x93, 0x68, 0xe2, 0xbc, 0x3a, 0x68, 0xde, 0x3b, 0x6b, 0xaf, 0xc3, 0x83,
	0x19, 0x45, 0x71, 0xc8, 0x68, 0x4c, 0xec, 0x87, 0xd0, 0x3a, 0xa3, 0xd1, 0x9c, 0x9e, 0xf6, 0x06,
	0xac, 0xcf, 0xcc, 0xaa, 0xe2, 0xcf, 0x50, 0x97, 0xea, 0xd5, 0x04, 0x33, 0xf7, 0x67, 0xfc, 0xc7,
	0xfe, 0x84, 0xa5, 0xc5, 0x6c, 0x7f, 0x8b, 0x60, 0xa6, 0xdd, 0xd5, 0x69, 0x3d, 0x68, 0x9c, 0x88,
	0x4f, 0x3d, 0x7f, 0xe4, 0x1f, 0xb7, 0x83, 0xf1, 0xb7, 0xdb, 0xc1, 0x6e, 0xc2, 0x72, 0xae, 0x85,
	0xea, 0xbb, 0x02, 0xe8, 0x8c, 0x86, 0x53, 0x9d, 0xed, 0x55, 0x68, 0xde, 0x42, 0x15, 0x79, 0x11,
	0x4c, 0x79, 0x15, 0x6a, 0xda, 0x12, 0x2c, 0xa8, 0x58, 0x11, 0xd6, 0x60, 0xe5, 0xd8, 0x8f, 0xf9,
	0x85, 0xbe, 0x08, 0x35, 0x71, 0x1f, 0x56, 0xa7, 0xf0, 0xb4, 0xe0, 0xf6, 0x85, 0x6a, 0xcc, 0xbd,
	0x50, 0xfb, 0x15, 0xf9, 0xa3, 0xb5, 0xfb, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xbb, 0x1a, 0x19, 0xab,
	0x26, 0x07, 0x00, 0x00,
}
