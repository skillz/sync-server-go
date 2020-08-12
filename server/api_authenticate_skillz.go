package server

import (
	"context"
	"github.com/aaron-skillz/sync-server-go/api"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ApiServerSkillz struct {
	*ApiServer
}

func (s *ApiServerSkillz) AuthenticateCustom(ctx context.Context, in *api.AuthenticateCustomRequest) (*api.Session, error) {
	s.logger.Debug("User ID", zap.String("user_id", in.Account.Id))
	// Before hook.
	if fn := s.runtime.BeforeAuthenticateCustom(); fn != nil {
		beforeFn := func(clientIP, clientPort string) error {
			result, err, code := fn(ctx, s.logger, "", "", nil, 0, clientIP, clientPort, in)
			if err != nil {
				return status.Error(code, err.Error())
			}
			if result == nil {
				// If result is nil, requested resource is disabled.
				s.logger.Warn("Intercepted a disabled resource.", zap.Any("resource", ctx.Value(ctxFullMethodKey{}).(string)))
				return status.Error(codes.NotFound, "Requested resource was not found.")
			}
			in = result
			return nil
		}

		// Execute the before function lambda wrapped in a trace for stats measurement.
		err := traceApiBefore(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), beforeFn)
		if err != nil {
			return nil, err
		}
	}

	var skToken string
	var ok bool
	if(in.Account.Vars == nil) {
		return nil, status.Error(codes.InvalidArgument, "Missing 'skillz_match_token' add dict to auth params.")
	} else if skToken, ok = in.Account.Vars["skillz_match_token"]; !ok || skToken == "" || !validateSkillzToken(skToken) {
		return nil, status.Error(codes.InvalidArgument, "Invalid skillz_match_token was sent.")
	}
	s.logger.Debug("Skillz token recived", zap.String("skillz_match_token", skToken))


	if in.Account == nil || in.Account.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Custom ID is required.")
	} else if invalidCharsRegex.MatchString(in.Account.Id) {
		return nil, status.Error(codes.InvalidArgument, "Custom ID invalid, no spaces or control characters allowed.")
	} else if len(in.Account.Id) < 6 || len(in.Account.Id) > 128 {
		return nil, status.Error(codes.InvalidArgument, "Custom ID invalid, must be 6-128 bytes.")
	}

	username := in.Username
	if username == "" {
		return nil, status.Error(codes.InvalidArgument, "Username must be set.")
	} else if invalidCharsRegex.MatchString(username) {
		return nil, status.Error(codes.InvalidArgument, "Username invalid, no spaces or control characters allowed.")
	} else if len(username) > 128 {
		return nil, status.Error(codes.InvalidArgument, "Username invalid, must be 1-128 bytes.")
	}

	token, exp := generateToken(s.config, in.Account.Id, username, in.Account.Vars)
	session := &api.Session{Created: false, Token: token}

	// After hook.
	if fn := s.runtime.AfterAuthenticateCustom(); fn != nil {
		afterFn := func(clientIP, clientPort string) error {
			return fn(ctx, s.logger, in.Account.Id, username, in.Account.Vars, exp, clientIP, clientPort, session, in)
		}

		// Execute the after function lambda wrapped in a trace for stats measurement.
		traceApiAfter(ctx, s.logger, s.metrics, ctx.Value(ctxFullMethodKey{}).(string), afterFn)
	}

	return session, nil
}

// TODO add token decryption once ready
func validateSkillzToken(token string) bool  {
	if len(token) < 10 {
		return false
	}
	return true
}

//func (aSkillz *ApiServerSkillz) AddFriends(ctx context.Context, request *api.AddFriendsRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) AddGroupUsers(ctx context.Context, request *api.AddGroupUsersRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) AuthenticateDevice(ctx context.Context, request *api.AuthenticateDeviceRequest) (*api.Session, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) AuthenticateEmail(ctx context.Context, request *api.AuthenticateEmailRequest) (*api.Session, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) BanGroupUsers(ctx context.Context, request *api.BanGroupUsersRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) BlockFriends(ctx context.Context, request *api.BlockFriendsRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) CreateGroup(ctx context.Context, request *api.CreateGroupRequest) (*api.Group, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) DeleteFriends(ctx context.Context, request *api.DeleteFriendsRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) DeleteGroup(ctx context.Context, request *api.DeleteGroupRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) DeleteLeaderboardRecord(ctx context.Context, request *api.DeleteLeaderboardRecordRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) DeleteNotifications(ctx context.Context, request *api.DeleteNotificationsRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) DeleteStorageObjects(ctx context.Context, request *api.DeleteStorageObjectsRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) Event(ctx context.Context, event *api.Event) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) GetAccount(ctx context.Context, empty *empty.Empty) (*api.Account, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) GetUsers(ctx context.Context, request *api.GetUsersRequest) (*api.Users, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) Healthcheck(ctx context.Context, empty *empty.Empty) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) JoinGroup(ctx context.Context, request *api.JoinGroupRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) JoinTournament(ctx context.Context, request *api.JoinTournamentRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) KickGroupUsers(ctx context.Context, request *api.KickGroupUsersRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) LeaveGroup(ctx context.Context, request *api.LeaveGroupRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) LinkCustom(ctx context.Context, custom *api.AccountCustom) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) LinkDevice(ctx context.Context, device *api.AccountDevice) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) LinkEmail(ctx context.Context, email *api.AccountEmail) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListChannelMessages(ctx context.Context, request *api.ListChannelMessagesRequest) (*api.ChannelMessageList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListFriends(ctx context.Context, request *api.ListFriendsRequest) (*api.FriendList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListGroups(ctx context.Context, request *api.ListGroupsRequest) (*api.GroupList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListGroupUsers(ctx context.Context, request *api.ListGroupUsersRequest) (*api.GroupUserList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListLeaderboardRecords(ctx context.Context, request *api.ListLeaderboardRecordsRequest) (*api.LeaderboardRecordList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListLeaderboardRecordsAroundOwner(ctx context.Context, request *api.ListLeaderboardRecordsAroundOwnerRequest) (*api.LeaderboardRecordList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListMatches(ctx context.Context, request *api.ListMatchesRequest) (*api.MatchList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListNotifications(ctx context.Context, request *api.ListNotificationsRequest) (*api.NotificationList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListStorageObjects(ctx context.Context, request *api.ListStorageObjectsRequest) (*api.StorageObjectList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListTournaments(ctx context.Context, request *api.ListTournamentsRequest) (*api.TournamentList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListTournamentRecords(ctx context.Context, request *api.ListTournamentRecordsRequest) (*api.TournamentRecordList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListTournamentRecordsAroundOwner(ctx context.Context, request *api.ListTournamentRecordsAroundOwnerRequest) (*api.TournamentRecordList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ListUserGroups(ctx context.Context, request *api.ListUserGroupsRequest) (*api.UserGroupList, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) PromoteGroupUsers(ctx context.Context, request *api.PromoteGroupUsersRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) ReadStorageObjects(ctx context.Context, request *api.ReadStorageObjectsRequest) (*api.StorageObjects, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) RpcFunc(ctx context.Context, rpc *api.Rpc) (*api.Rpc, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) UnlinkCustom(ctx context.Context, custom *api.AccountCustom) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) UnlinkDevice(ctx context.Context, device *api.AccountDevice) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) UnlinkEmail(ctx context.Context, email *api.AccountEmail) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) UpdateAccount(ctx context.Context, request *api.UpdateAccountRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) UpdateGroup(ctx context.Context, request *api.UpdateGroupRequest) (*empty.Empty, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) WriteLeaderboardRecord(ctx context.Context, request *api.WriteLeaderboardRecordRequest) (*api.LeaderboardRecord, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) WriteStorageObjects(ctx context.Context, request *api.WriteStorageObjectsRequest) (*api.StorageObjectAcks, error) {
//	panic("implement me")
//}
//
//func (aSkillz *ApiServerSkillz) WriteTournamentRecord(ctx context.Context, request *api.WriteTournamentRecordRequest) (*api.LeaderboardRecord, error) {
//	panic("implement me")
//}
