package service

import (
	"context"
	"github.com/kirpepa/spendly-api/group/model"
	"github.com/kirpepa/spendly-api/group/proto"
	"github.com/kirpepa/spendly-api/group/repository"
)

type GroupServer struct {
	proto.UnimplementedGroupServiceServer
	repo *repository.GroupRepo
}

func NewGroupServer(repo *repository.GroupRepo) *GroupServer {
	return &GroupServer{repo: repo}
}

func (s *GroupServer) CreateGroup(ctx context.Context, req *proto.CreateGroupRequest) (*proto.GroupResponse, error) {
	group := &model.Group{
		Name:        req.Name,
		Description: req.Description,
		OwnerID:     uint(req.OwnerId),
	}
	if err := s.repo.Create(group); err != nil {
		return nil, err
	}
	return &proto.GroupResponse{
		Id:          uint64(group.ID),
		Name:        group.Name,
		Description: group.Description,
		OwnerId:     uint64(group.OwnerID),
	}, nil
}

func (s *GroupServer) GetGroup(ctx context.Context, req *proto.GetGroupRequest) (*proto.GroupResponse, error) {
	group, err := s.repo.GetByID(uint(req.GroupId))
	if err != nil {
		return nil, err
	}
	return &proto.GroupResponse{
		Id:          uint64(group.ID),
		Name:        group.Name,
		Description: group.Description,
		OwnerId:     uint64(group.OwnerID),
	}, nil
}

func (s *GroupServer) ListGroups(ctx context.Context, req *proto.ListGroupsRequest) (*proto.ListGroupsResponse, error) {
	groups, err := s.repo.ListByOwner(uint(req.OwnerId))
	if err != nil {
		return nil, err
	}

	var protoGroups []*proto.GroupResponse
	for _, g := range groups {
		protoGroups = append(protoGroups, &proto.GroupResponse{
			Id:          uint64(g.ID),
			Name:        g.Name,
			Description: g.Description,
			OwnerId:     uint64(g.OwnerID),
		})
	}

	return &proto.ListGroupsResponse{Groups: protoGroups}, nil
}
