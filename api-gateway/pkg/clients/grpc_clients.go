package clients

import (
	"context"
	"time"

	authpb "github.com/kirpepa/spendly-api/auth/proto"
	expbp "github.com/kirpepa/spendly-api/expense/proto"
	grouppb "github.com/kirpepa/spendly-api/group/proto"
	gmembpb "github.com/kirpepa/spendly-api/group_member/proto"
	userpb "github.com/kirpepa/spendly-api/user/proto"

	"google.golang.org/grpc"
)

type GRPCClients struct {
	AuthClient        authpb.AuthServiceClient
	UserClient        userpb.UserServiceClient
	GroupClient       grouppb.GroupServiceClient
	GroupMemberClient gmembpb.GroupMemberServiceClient
	ExpenseClient     expbp.ExpenseServiceClient

	authConn, userConn, groupConn, gmConn, expConn *grpc.ClientConn
}

func Dial(addrs map[string]string, timeout time.Duration) (*GRPCClients, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	var err error
	c := &GRPCClients{}

	c.authConn, err = grpc.DialContext(ctx, addrs["auth"], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.AuthClient = authpb.NewAuthServiceClient(c.authConn)

	c.userConn, err = grpc.DialContext(ctx, addrs["user"], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.UserClient = userpb.NewUserServiceClient(c.userConn)

	c.groupConn, err = grpc.DialContext(ctx, addrs["group"], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.GroupClient = grouppb.NewGroupServiceClient(c.groupConn)

	c.gmConn, err = grpc.DialContext(ctx, addrs["group_member"], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.GroupMemberClient = gmembpb.NewGroupMemberServiceClient(c.gmConn)

	c.expConn, err = grpc.DialContext(ctx, addrs["expense"], grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	c.ExpenseClient = expbp.NewExpenseServiceClient(c.expConn)

	return c, nil
}

func (c *GRPCClients) Close() {
	if c.authConn != nil {
		_ = c.authConn.Close()
	}
	if c.userConn != nil {
		_ = c.userConn.Close()
	}
	if c.groupConn != nil {
		_ = c.groupConn.Close()
	}
	if c.gmConn != nil {
		_ = c.gmConn.Close()
	}
	if c.expConn != nil {
		_ = c.expConn.Close()
	}

}
