package lp_handlers

import (
	"context"
	"errors"
	"time"

	chanserv "github.com/DimTur/lp_learning_platform/internal/services/channel"
	"github.com/DimTur/lp_learning_platform/internal/services/storage/postgresql/channels"
	lpv1 "github.com/DimTur/lp_protos/gen/go/lp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func (s *serverAPI) CreateChannel(ctx context.Context, req *lpv1.CreateChannelRequest) (*lpv1.CreateChannelResponse, error) {
	channel := channels.CreateChannel{
		Name:            req.GetName(),
		Description:     req.GetDescription(),
		CreatedBy:       req.GetCreatedBy(),
		LastModifiedBy:  req.GetCreatedBy(),
		LearningGroupId: req.GetLearningGroupId(),
	}
	channelID, err := s.channelHandlers.CreateChannel(ctx, &channel)
	if err != nil {
		if errors.Is(err, chanserv.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &lpv1.CreateChannelResponse{
		Id: channelID,
	}, nil
}

func (s *serverAPI) GetChannel(ctx context.Context, req *lpv1.GetChannelRequest) (*lpv1.GetChannelResponse, error) {
	channel, err := s.channelHandlers.GetChannel(ctx, &channels.GetChannelByID{
		ChannelID: req.GetChannelId(),
	})
	if err != nil {
		if errors.Is(err, chanserv.ErrChannelNotFound) {
			return nil, status.Error(codes.NotFound, "channel not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	var plans []*lpv1.Plan
	for _, plan := range channel.Plans {
		plans = append(plans, &lpv1.Plan{
			Id:             plan.ID,
			Name:           plan.Name,
			Description:    plan.Description,
			CreatedBy:      plan.CreatedBy,
			LastModifiedBy: plan.LastModifiedBy,
			IsPublished:    plan.IsPublished,
			Public:         plan.Public,
			CreatedAt:      plan.CreatedAt.Format(time.RFC3339),
			Modified:       plan.Modified.Format(time.RFC3339),
		})
	}

	return &lpv1.GetChannelResponse{
		Channel: &lpv1.ChannelWithPlans{
			Id:             channel.ID,
			Name:           channel.Name,
			Description:    channel.Description,
			CreatedBy:      channel.CreatedBy,
			LastModifiedBy: channel.LastModifiedBy,
			CreatedAt:      channel.CreatedAt.Format(time.RFC3339),
			Modified:       channel.Modified.Format(time.RFC3339),
			Plans:          plans,
		},
	}, nil
}

func (s *serverAPI) GetChannels(ctx context.Context, req *lpv1.GetChannelsRequest) (*lpv1.GetChannelsResponse, error) {
	channels, err := s.channelHandlers.GetChannels(ctx, &channels.GetChannels{
		LgIDs:  req.GetLearningGroupIds(),
		Limit:  req.GetLimit(),
		Offset: req.GetOffset(),
	})
	if err != nil {
		switch {
		case errors.Is(err, chanserv.ErrChannelNotFound):
			return nil, status.Error(codes.NotFound, "channels not found")
		case errors.Is(err, chanserv.ErrInvalidCredentials):
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	var responseChannels []*lpv1.Channel
	for _, channel := range channels {
		responseChannels = append(responseChannels, &lpv1.Channel{
			Id:             channel.ID,
			Name:           channel.Name,
			Description:    channel.Description,
			CreatedBy:      channel.CreatedBy,
			LastModifiedBy: channel.LastModifiedBy,
			CreatedAt:      channel.CreatedAt.Format(time.RFC3339),
			Modified:       channel.Modified.Format(time.RFC3339),
		})
	}

	return &lpv1.GetChannelsResponse{
		Channels: responseChannels,
	}, nil
}

func (s *serverAPI) UpdateChannel(ctx context.Context, req *lpv1.UpdateChannelRequest) (*lpv1.UpdateChannelResponse, error) {
	var name *string
	if req.GetName() != "" {
		name = proto.String(req.GetName())
	}

	var description *string
	if req.GetDescription() != "" {
		description = proto.String(req.GetDescription())
	}

	updChannel := channels.UpdateChannelRequest{
		UserID:      req.GetUserId(),
		ChannelID:   req.GetChannelId(),
		Name:        name,
		Description: description,
	}

	id, err := s.channelHandlers.UpdateChannel(ctx, &updChannel)
	if err != nil {
		switch {
		case errors.Is(err, chanserv.ErrInvalidCredentials):
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &lpv1.UpdateChannelResponse{
		Id: id,
	}, nil
}

func (s *serverAPI) DeleteChannel(ctx context.Context, req *lpv1.DeleteChannelRequest) (*lpv1.DeleteChannelResponse, error) {
	err := s.channelHandlers.DeleteChannel(ctx, &channels.DeleteChannelRequest{
		ChannelID: req.GetChannelId(),
	})
	if err != nil {
		if errors.Is(err, chanserv.ErrChannelNotFound) {
			return nil, status.Error(codes.NotFound, "channel not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &lpv1.DeleteChannelResponse{
		Success: true,
	}, nil
}

func (s *serverAPI) ShareChannelToGroup(ctx context.Context, req *lpv1.ShareChannelToGroupRequest) (*lpv1.ShareChannelToGroupResponse, error) {
	sharingChannel := channels.ShareChannelToGroup{
		ChannelID: req.GetChannelId(),
		LGroupIDs: req.GetLgroupsIds(),
		CreatedBy: req.GetCreatedBy(),
	}
	if err := s.channelHandlers.ShareChannelToGroup(ctx, sharingChannel); err != nil {
		if errors.Is(err, chanserv.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &lpv1.ShareChannelToGroupResponse{
		Success: true,
	}, nil
}

func (s *serverAPI) IsChannelCreator(ctx context.Context, req *lpv1.IsChannelCreatorRequest) (*lpv1.IsChannelCreatorResponse, error) {

	isCreator, err := s.channelHandlers.IsChannelCreator(ctx, &channels.IsChannelCreator{
		UserID:    req.GetUserId(),
		ChannelID: req.GetChannelId(),
	})
	if err != nil {
		switch {
		case errors.Is(err, chanserv.ErrChannelNotFound):
			return nil, status.Error(codes.NotFound, "channel not found")
		case errors.Is(err, chanserv.ErrPermissionDenied):
			return nil, status.Error(codes.PermissionDenied, "permission denied")
		case errors.Is(err, chanserv.ErrInvalidCredentials):
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &lpv1.IsChannelCreatorResponse{
		IsCreator: isCreator,
	}, nil
}

func (s *serverAPI) GetLearningGroupsShareWithChannel(ctx context.Context, req *lpv1.GetLearningGroupsShareWithChannelRequest) (*lpv1.GetLearningGroupsShareWithChannelResponse, error) {
	lgIDs, err := s.channelHandlers.GetLearningGroupsShareWithChannel(ctx, req.ChannelId)
	if err != nil {
		switch {
		case errors.Is(err, chanserv.ErrInvalidCredentials):
			return nil, status.Error(codes.InvalidArgument, "bad request")
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	return &lpv1.GetLearningGroupsShareWithChannelResponse{
		LearningGroupIds: lgIDs,
	}, nil
}
