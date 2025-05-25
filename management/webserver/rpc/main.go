package rpc

import (
	"net"
	"time"

	"google.golang.org/grpc"

	"xianfengyuan.cn/dev/go/errors"
	"xianfengyuan.cn/dev/go/log"
	"xianfengyuan.cn/patronus/safeline-2/management/webserver/pkg/config"
	pb "xianfengyuan.cn/patronus/safeline-2/management/webserver/proto/website"
)

const (
	WaitRspTimeout   = 30 * time.Second
	KeepaliveTime    = 5 * time.Second
	KeepaliveTimeout = 30 * time.Second

	Ping                   = "ping"
	Pong                   = "pong"
	EventTypeWebsite       = "website"
	EventTypeDeleteWebsite = "deleteWebsite"
	EventTypeFullWebsite   = "fullWebsite"
)

var logger = log.GetLogger("grpc")

func StartGRPCSever() error {
	lis, err := net.Listen("tcp", config.GlobalConfig.GPRC.ListenAddr)
	if err != nil {
		return errors.Wrap(err, "Failed to listen")
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterWebsiteServer(grpcServer, GetWebsiteServer())
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			logger.Fatalln("Failed to server")
		}
	}()
	return nil
}
