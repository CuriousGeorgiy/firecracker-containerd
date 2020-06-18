// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: fccontrol.proto

package fccontrol

import (
	context "context"
	fmt "fmt"
	github_com_containerd_ttrpc "github.com/containerd/ttrpc"
	proto1 "github.com/firecracker-microvm/firecracker-containerd/proto"
	proto "github.com/gogo/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("fccontrol.proto", fileDescriptor_b99f53e2bf82c5ef) }

var fileDescriptor_b99f53e2bf82c5ef = []byte{
<<<<<<< HEAD
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x3f, 0x4f, 0xf2, 0x50,
	0x14, 0xc6, 0xe9, 0xf0, 0xf2, 0xc2, 0x35, 0x08, 0xdc, 0x04, 0x45, 0x4c, 0xba, 0xb8, 0x1f, 0x0c,
	0x3a, 0x33, 0x80, 0x5a, 0x1d, 0x9a, 0x18, 0x88, 0x0c, 0x6e, 0x97, 0x72, 0x4a, 0x88, 0xa5, 0xb7,
	0xb6, 0xa7, 0x83, 0x9b, 0x1f, 0x8f, 0xd1, 0xd1, 0x51, 0xba, 0xfb, 0x1d, 0x0c, 0xb4, 0x97, 0x3f,
	0x85, 0xe4, 0x6e, 0xcf, 0xf9, 0x9d, 0x9e, 0xe7, 0x3c, 0xf7, 0x24, 0x65, 0x55, 0xd7, 0x71, 0xa4,
	0x4f, 0xa1, 0xf4, 0x20, 0x08, 0x25, 0xc9, 0xd6, 0xe5, 0x54, 0xca, 0xa9, 0x87, 0xed, 0x75, 0x35,
	0x8e, 0xdd, 0x36, 0xce, 0x03, 0xfa, 0xc8, 0x9a, 0x75, 0x77, 0x16, 0xa2, 0x13, 0x0a, 0xe7, 0x0d,
	0xc3, 0x14, 0x75, 0x7e, 0xff, 0xb1, 0x93, 0x87, 0x2d, 0xe5, 0x6d, 0x56, 0xea, 0x87, 0x28, 0x08,
	0x47, 0x36, 0xaf, 0x81, 0x92, 0x03, 0x7c, 0x8f, 0x31, 0xa2, 0x56, 0x7d, 0x87, 0x44, 0x81, 0xf4,
	0x23, 0xe4, 0x1d, 0xf6, 0xff, 0x59, 0xc4, 0xd1, 0xea, 0xfb, 0x2a, 0x64, 0x4a, 0x7d, 0x7e, 0x06,
	0x69, 0x1a, 0x50, 0x69, 0xe0, 0x7e, 0x95, 0x86, 0xdf, 0xb2, 0xd2, 0x00, 0xa3, 0x78, 0x9e, 0x2e,
	0x51, 0x52, 0x37, 0x75, 0xcd, 0x8a, 0x43, 0x92, 0xc1, 0xc8, 0xe6, 0xa7, 0x90, 0x0a, 0xdd, 0x44,
	0x87, 0x95, 0x2d, 0xa4, 0x91, 0xfd, 0xe4, 0xbb, 0x92, 0xd7, 0x61, 0xa3, 0xd5, 0x1c, 0xdf, 0x45,
	0xd9, 0x7b, 0xba, 0xac, 0x32, 0x5c, 0x41, 0x1b, 0x49, 0x4c, 0x04, 0x09, 0xde, 0x80, 0xbd, 0x5a,
	0xb7, 0xf3, 0x8e, 0xd5, 0x5e, 0x82, 0xc9, 0xfa, 0x46, 0x1b, 0x8b, 0x26, 0xe4, 0x91, 0xce, 0xa5,
	0xcb, 0x2a, 0x56, 0x2e, 0x85, 0x75, 0x3c, 0x45, 0x0e, 0x67, 0xaf, 0xb0, 0x58, 0xcd, 0x42, 0xea,
	0x09, 0xcf, 0x93, 0xd2, 0xef, 0x4b, 0xdf, 0x9d, 0x4d, 0x79, 0x13, 0xf2, 0x48, 0xb9, 0x5c, 0x1c,
	0xe9, 0x6c, 0xcf, 0x91, 0x66, 0xcf, 0xda, 0xbc, 0x01, 0x7b, 0xb5, 0xfe, 0x1c, 0xd5, 0xad, 0xf7,
	0x90, 0x04, 0x45, 0xfc, 0x1c, 0x72, 0x44, 0x79, 0x34, 0x0f, 0x1b, 0x59, 0x8a, 0x47, 0xc6, 0xf7,
	0xb6, 0xa6, 0x46, 0x2d, 0x38, 0x84, 0x9a, 0x3c, 0xbd, 0xab, 0xc5, 0xd2, 0x2c, 0x7c, 0x2f, 0xcd,
	0xc2, 0x67, 0x62, 0x1a, 0x8b, 0xc4, 0x34, 0xbe, 0x12, 0xd3, 0xf8, 0x49, 0x4c, 0xe3, 0xb5, 0xbc,
	0xf9, 0x95, 0xc6, 0xc5, 0xf5, 0xd0, 0xcd, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x66, 0x43, 0x0c,
	0x39, 0x5e, 0x03, 0x00, 0x00,
=======
<<<<<<< HEAD
	// 290 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4f, 0x4b, 0x4e, 0xce,
	0xcf, 0x2b, 0x29, 0xca, 0xcf, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x4e, 0xcf, 0xcf,
	0x4f, 0xcf, 0x49, 0xd5, 0x07, 0xf3, 0x92, 0x4a, 0xd3, 0xf4, 0x53, 0x73, 0x0b, 0x4a, 0x2a, 0xa1,
	0x92, 0x82, 0x69, 0x99, 0x45, 0xa9, 0xc9, 0x45, 0x89, 0xc9, 0xd9, 0xa9, 0x45, 0x10, 0x21, 0xa3,
	0x2b, 0xcc, 0x5c, 0xdc, 0x6e, 0x08, 0x51, 0x21, 0x7d, 0x2e, 0x0e, 0xe7, 0xa2, 0xd4, 0xc4, 0x92,
	0xd4, 0x30, 0x5f, 0x21, 0x01, 0x3d, 0x18, 0x33, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x4a,
	0x10, 0x49, 0xa4, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55, 0xc8, 0x88, 0x8b, 0x3d, 0x20, 0xb1, 0xb4,
	0x18, 0xa4, 0x9e, 0x5f, 0x0f, 0xca, 0x82, 0x29, 0x17, 0xd3, 0x83, 0xb8, 0x46, 0x0f, 0xe6, 0x1a,
	0x3d, 0x57, 0x90, 0x6b, 0x84, 0x4c, 0xb8, 0x38, 0x82, 0x52, 0x8b, 0x4b, 0x73, 0x21, 0x96, 0xc0,
	0x98, 0x84, 0x74, 0x19, 0x70, 0xb1, 0x05, 0x97, 0xe4, 0x17, 0x84, 0xf9, 0x0a, 0xf1, 0xe9, 0x41,
	0x18, 0x84, 0x74, 0x18, 0x71, 0x71, 0xba, 0xa7, 0x96, 0x84, 0xf9, 0x7a, 0xe6, 0xa5, 0xe5, 0x0b,
	0x09, 0xea, 0xc1, 0xd9, 0x30, 0x7d, 0x42, 0xc8, 0x42, 0x50, 0xff, 0xd8, 0x71, 0xf1, 0x06, 0x83,
	0x04, 0x7d, 0x53, 0x4b, 0x12, 0x53, 0x12, 0x4b, 0x12, 0x85, 0x44, 0xf5, 0x50, 0xf8, 0x84, 0xec,
	0x74, 0xe1, 0x12, 0x08, 0x2d, 0x48, 0x01, 0x87, 0x11, 0xdc, 0x08, 0x09, 0x3d, 0x74, 0x21, 0x42,
	0xa6, 0xd8, 0x71, 0xf1, 0xba, 0xa3, 0xb9, 0xc2, 0x1d, 0xbb, 0x2b, 0xd0, 0x84, 0x21, 0xbe, 0x70,
	0x52, 0x3e, 0xf1, 0x50, 0x8e, 0xe1, 0xc6, 0x43, 0x39, 0x86, 0x86, 0x47, 0x72, 0x8c, 0x27, 0x1e,
	0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0x63, 0x14, 0x27, 0x3c, 0xc5, 0x24,
	0xb1, 0x81, 0x2d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0x26, 0xe6, 0x13, 0x45, 0x02,
	0x00, 0x00,
=======
	// 336 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xbb, 0x4e, 0xf3, 0x40,
	0x10, 0x85, 0xe3, 0xe2, 0xcf, 0x65, 0x7e, 0x72, 0x5b, 0x41, 0x84, 0x82, 0xe4, 0x86, 0x7e, 0x82,
	0x02, 0x25, 0x8a, 0x10, 0xb7, 0x08, 0x09, 0x0b, 0x94, 0x88, 0x14, 0x74, 0x1b, 0x7b, 0x1c, 0x10,
	0x89, 0xc7, 0xd8, 0xeb, 0x82, 0x8e, 0x92, 0x47, 0x4b, 0x49, 0x49, 0x49, 0xfc, 0x24, 0x28, 0xf1,
	0x85, 0x24, 0x42, 0xda, 0xee, 0xcc, 0x67, 0x9f, 0x99, 0xb3, 0x3b, 0x0b, 0x75, 0xd7, 0xb6, 0xd9,
	0x53, 0x01, 0x4f, 0xd1, 0x0f, 0x58, 0x71, 0xfb, 0x60, 0xc2, 0x3c, 0x99, 0x52, 0x67, 0x55, 0x8d,
	0x23, 0xb7, 0x43, 0x33, 0x5f, 0xbd, 0xa5, 0x1f, 0x9b, 0xee, 0x73, 0x40, 0x76, 0x20, 0xed, 0x17,
	0x0a, 0x12, 0xd4, 0xfd, 0xf8, 0x07, 0xff, 0xaf, 0x7f, 0xa9, 0xe8, 0x40, 0xf9, 0x22, 0x20, 0xa9,
	0x68, 0x64, 0x89, 0x06, 0x66, 0x72, 0x40, 0xaf, 0x11, 0x85, 0xaa, 0xdd, 0x5c, 0x23, 0xa1, 0xcf,
	0x5e, 0x48, 0xe2, 0x08, 0x8a, 0x43, 0xc5, 0xfe, 0xc8, 0x12, 0x35, 0x4c, 0x44, 0xf6, 0x73, 0x0b,
	0x93, 0x2c, 0x98, 0x65, 0xc1, 0xab, 0x65, 0x16, 0xd1, 0x85, 0x4a, 0x9f, 0xd4, 0xc8, 0xba, 0xf1,
	0x5c, 0x16, 0x4d, 0xcc, 0x75, 0xe6, 0x13, 0xeb, 0x28, 0x9d, 0xd2, 0x83, 0xea, 0x70, 0x09, 0x2d,
	0x52, 0xd2, 0x91, 0x4a, 0x8a, 0x3d, 0xdc, 0xa8, 0x75, 0x33, 0x2f, 0xa1, 0xf1, 0xe0, 0x3b, 0xab,
	0xe4, 0x79, 0x8b, 0x7d, 0xdc, 0x46, 0xba, 0x2e, 0x3d, 0xa8, 0xf6, 0xb7, 0x52, 0xf4, 0xff, 0x4e,
	0xb1, 0x85, 0xd3, 0x53, 0x74, 0xa1, 0x74, 0x2f, 0xa3, 0x70, 0x79, 0xb7, 0x75, 0x4c, 0x95, 0x6e,
	0xe6, 0x09, 0x94, 0x07, 0x14, 0x46, 0xb3, 0x64, 0x21, 0x99, 0xd4, 0xb9, 0x4e, 0x61, 0xe7, 0x96,
	0xa5, 0x33, 0xf4, 0xa4, 0x1f, 0x3e, 0xb1, 0x12, 0xbb, 0xb8, 0x5e, 0xea, 0xdc, 0x67, 0x50, 0x4b,
	0xf6, 0x9c, 0xfb, 0x5b, 0xb8, 0x09, 0xf4, 0x3b, 0x2e, 0xdd, 0xb9, 0xee, 0x94, 0xa5, 0x23, 0xea,
	0x98, 0x2a, 0x8d, 0xe7, 0xfc, 0x70, 0xbe, 0x30, 0x0b, 0x5f, 0x0b, 0xb3, 0xf0, 0x1e, 0x9b, 0xc6,
	0x3c, 0x36, 0x8d, 0xcf, 0xd8, 0x34, 0xbe, 0x63, 0xd3, 0x78, 0xac, 0xe4, 0xaf, 0x7c, 0x5c, 0x5c,
	0x99, 0x8e, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfb, 0xe9, 0x16, 0xe0, 0xf9, 0x02, 0x00, 0x00,
>>>>>>> Added support for creating and loading snapshots of VM.
>>>>>>> Added support for creating and loading snapshots of VM.
}

type FirecrackerService interface {
	CreateVM(ctx context.Context, req *proto1.CreateVMRequest) (*proto1.CreateVMResponse, error)
	PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error)
	ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error)
	StopVM(ctx context.Context, req *proto1.StopVMRequest) (*empty.Empty, error)
	GetVMInfo(ctx context.Context, req *proto1.GetVMInfoRequest) (*proto1.GetVMInfoResponse, error)
	SetVMMetadata(ctx context.Context, req *proto1.SetVMMetadataRequest) (*empty.Empty, error)
	UpdateVMMetadata(ctx context.Context, req *proto1.UpdateVMMetadataRequest) (*empty.Empty, error)
	GetVMMetadata(ctx context.Context, req *proto1.GetVMMetadataRequest) (*proto1.GetVMMetadataResponse, error)
<<<<<<< HEAD
	GetBalloonConfig(ctx context.Context, req *proto1.GetBalloonConfigRequest) (*proto1.GetBalloonConfigResponse, error)
	UpdateBalloon(ctx context.Context, req *proto1.UpdateBalloonRequest) (*empty.Empty, error)
	GetBalloonStats(ctx context.Context, req *proto1.GetBalloonStatsRequest) (*proto1.GetBalloonStatsResponse, error)
	UpdateBalloonStats(ctx context.Context, req *proto1.UpdateBalloonStatsRequest) (*empty.Empty, error)
=======
<<<<<<< HEAD
=======
	PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error)
	ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error)
	LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*empty.Empty, error)
	CreateSnapshot(ctx context.Context, req *proto1.CreateSnapshotRequest) (*empty.Empty, error)
	Offload(ctx context.Context, req *proto1.OffloadRequest) (*empty.Empty, error)
>>>>>>> Added support for creating and loading snapshots of VM.
>>>>>>> Added support for creating and loading snapshots of VM.
}

func RegisterFirecrackerService(srv *github_com_containerd_ttrpc.Server, svc FirecrackerService) {
	srv.Register("Firecracker", map[string]github_com_containerd_ttrpc.Method{
		"CreateVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.CreateVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.CreateVM(ctx, &req)
		},
		"PauseVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.PauseVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.PauseVM(ctx, &req)
		},
		"ResumeVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.ResumeVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.ResumeVM(ctx, &req)
		},
		"StopVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.StopVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.StopVM(ctx, &req)
		},
		"GetVMInfo": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetVMInfoRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetVMInfo(ctx, &req)
		},
		"SetVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.SetVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.SetVMMetadata(ctx, &req)
		},
		"UpdateVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.UpdateVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.UpdateVMMetadata(ctx, &req)
		},
		"GetVMMetadata": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetVMMetadataRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetVMMetadata(ctx, &req)
		},
<<<<<<< HEAD
		"GetBalloonConfig": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetBalloonConfigRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetBalloonConfig(ctx, &req)
		},
		"UpdateBalloon": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.UpdateBalloonRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.UpdateBalloon(ctx, &req)
		},
		"GetBalloonStats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.GetBalloonStatsRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.GetBalloonStats(ctx, &req)
		},
		"UpdateBalloonStats": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.UpdateBalloonStatsRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.UpdateBalloonStats(ctx, &req)
		},
=======
<<<<<<< HEAD
=======
		"PauseVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.PauseVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.PauseVM(ctx, &req)
		},
		"ResumeVM": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.ResumeVMRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.ResumeVM(ctx, &req)
		},
		"LoadSnapshot": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.LoadSnapshotRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.LoadSnapshot(ctx, &req)
		},
		"CreateSnapshot": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.CreateSnapshotRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.CreateSnapshot(ctx, &req)
		},
		"Offload": func(ctx context.Context, unmarshal func(interface{}) error) (interface{}, error) {
			var req proto1.OffloadRequest
			if err := unmarshal(&req); err != nil {
				return nil, err
			}
			return svc.Offload(ctx, &req)
		},
>>>>>>> Added support for creating and loading snapshots of VM.
>>>>>>> Added support for creating and loading snapshots of VM.
	})
}

type firecrackerClient struct {
	client *github_com_containerd_ttrpc.Client
}

func NewFirecrackerClient(client *github_com_containerd_ttrpc.Client) FirecrackerService {
	return &firecrackerClient{
		client: client,
	}
}

func (c *firecrackerClient) CreateVM(ctx context.Context, req *proto1.CreateVMRequest) (*proto1.CreateVMResponse, error) {
	var resp proto1.CreateVMResponse
	if err := c.client.Call(ctx, "Firecracker", "CreateVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "PauseVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "ResumeVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) StopVM(ctx context.Context, req *proto1.StopVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "StopVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) GetVMInfo(ctx context.Context, req *proto1.GetVMInfoRequest) (*proto1.GetVMInfoResponse, error) {
	var resp proto1.GetVMInfoResponse
	if err := c.client.Call(ctx, "Firecracker", "GetVMInfo", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) SetVMMetadata(ctx context.Context, req *proto1.SetVMMetadataRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "SetVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) UpdateVMMetadata(ctx context.Context, req *proto1.UpdateVMMetadataRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "UpdateVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) GetVMMetadata(ctx context.Context, req *proto1.GetVMMetadataRequest) (*proto1.GetVMMetadataResponse, error) {
	var resp proto1.GetVMMetadataResponse
	if err := c.client.Call(ctx, "Firecracker", "GetVMMetadata", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
<<<<<<< HEAD

func (c *firecrackerClient) GetBalloonConfig(ctx context.Context, req *proto1.GetBalloonConfigRequest) (*proto1.GetBalloonConfigResponse, error) {
	var resp proto1.GetBalloonConfigResponse
	if err := c.client.Call(ctx, "Firecracker", "GetBalloonConfig", req, &resp); err != nil {
=======
<<<<<<< HEAD
=======

func (c *firecrackerClient) PauseVM(ctx context.Context, req *proto1.PauseVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "PauseVM", req, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *firecrackerClient) ResumeVM(ctx context.Context, req *proto1.ResumeVMRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "ResumeVM", req, &resp); err != nil {
>>>>>>> Added support for creating and loading snapshots of VM.
		return nil, err
	}
	return &resp, nil
}

<<<<<<< HEAD
func (c *firecrackerClient) UpdateBalloon(ctx context.Context, req *proto1.UpdateBalloonRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "UpdateBalloon", req, &resp); err != nil {
=======
func (c *firecrackerClient) LoadSnapshot(ctx context.Context, req *proto1.LoadSnapshotRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "LoadSnapshot", req, &resp); err != nil {
>>>>>>> Added support for creating and loading snapshots of VM.
		return nil, err
	}
	return &resp, nil
}

<<<<<<< HEAD
func (c *firecrackerClient) GetBalloonStats(ctx context.Context, req *proto1.GetBalloonStatsRequest) (*proto1.GetBalloonStatsResponse, error) {
	var resp proto1.GetBalloonStatsResponse
	if err := c.client.Call(ctx, "Firecracker", "GetBalloonStats", req, &resp); err != nil {
=======
func (c *firecrackerClient) CreateSnapshot(ctx context.Context, req *proto1.CreateSnapshotRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "CreateSnapshot", req, &resp); err != nil {
>>>>>>> Added support for creating and loading snapshots of VM.
		return nil, err
	}
	return &resp, nil
}

<<<<<<< HEAD
func (c *firecrackerClient) UpdateBalloonStats(ctx context.Context, req *proto1.UpdateBalloonStatsRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "UpdateBalloonStats", req, &resp); err != nil {
=======
func (c *firecrackerClient) Offload(ctx context.Context, req *proto1.OffloadRequest) (*empty.Empty, error) {
	var resp empty.Empty
	if err := c.client.Call(ctx, "Firecracker", "Offload", req, &resp); err != nil {
>>>>>>> Added support for creating and loading snapshots of VM.
		return nil, err
	}
	return &resp, nil
}
<<<<<<< HEAD
=======
>>>>>>> Added support for creating and loading snapshots of VM.
>>>>>>> Added support for creating and loading snapshots of VM.
