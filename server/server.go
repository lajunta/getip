package main

import (
	"context"
	"log"

	"github.com/lajunta/getip/grpcd"
)

type GetIPServer struct {
	grpcd.UnimplementedGetIPServiceServer
}

func (s *GetIPServer) GetIP(ctx context.Context, in *grpcd.IPRequest) (*grpcd.IPReply, error) {
	log.Printf("room %s mac %s requesting .", in.Room, in.Mac)
	pc := rooms[in.Room][in.Mac]
	hostinfo := &grpcd.IPReply{
		Mac:       in.Mac,
		Name:      pc.Name,
		WorkGroup: pc.WorkGroup,
		IP:        pc.IP,
		NetMask:   pc.Netmask,
		GateWay:   pc.Gateway,
		DNS:       pc.DNS,
	}

	return hostinfo, nil
}
