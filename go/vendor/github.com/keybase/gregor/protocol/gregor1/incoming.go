// Auto-generated by avdl-compiler v1.3.3 (https://github.com/keybase/node-avdl-compiler)
//   Input file: gregor1/avdl/incoming.avdl

package gregor1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
	context "golang.org/x/net/context"
)

type SyncResult struct {
	Msgs []InBandMessage `codec:"msgs" json:"msgs"`
	Hash []byte          `codec:"hash" json:"hash"`
}

type SyncArg struct {
	Uid      UID      `codec:"uid" json:"uid"`
	Deviceid DeviceID `codec:"deviceid" json:"deviceid"`
	Ctime    Time     `codec:"ctime" json:"ctime"`
}

type ConsumeMessageArg struct {
	M Message `codec:"m" json:"m"`
}

type ConsumePublishMessageArg struct {
	M Message `codec:"m" json:"m"`
}

type PingArg struct {
}

type IncomingInterface interface {
	Sync(context.Context, SyncArg) (SyncResult, error)
	ConsumeMessage(context.Context, Message) error
	ConsumePublishMessage(context.Context, Message) error
	Ping(context.Context) (string, error)
}

func IncomingProtocol(i IncomingInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "gregor.1.incoming",
		Methods: map[string]rpc.ServeHandlerDescription{
			"sync": {
				MakeArg: func() interface{} {
					ret := make([]SyncArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]SyncArg)
					if !ok {
						err = rpc.NewTypeError((*[]SyncArg)(nil), args)
						return
					}
					ret, err = i.Sync(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"consumeMessage": {
				MakeArg: func() interface{} {
					ret := make([]ConsumeMessageArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConsumeMessageArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConsumeMessageArg)(nil), args)
						return
					}
					err = i.ConsumeMessage(ctx, (*typedArgs)[0].M)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"consumePublishMessage": {
				MakeArg: func() interface{} {
					ret := make([]ConsumePublishMessageArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]ConsumePublishMessageArg)
					if !ok {
						err = rpc.NewTypeError((*[]ConsumePublishMessageArg)(nil), args)
						return
					}
					err = i.ConsumePublishMessage(ctx, (*typedArgs)[0].M)
					return
				},
				MethodType: rpc.MethodCall,
			},
			"ping": {
				MakeArg: func() interface{} {
					ret := make([]PingArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					ret, err = i.Ping(ctx)
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type IncomingClient struct {
	Cli rpc.GenericClient
}

func (c IncomingClient) Sync(ctx context.Context, __arg SyncArg) (res SyncResult, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.sync", []interface{}{__arg}, &res)
	return
}

func (c IncomingClient) ConsumeMessage(ctx context.Context, m Message) (err error) {
	__arg := ConsumeMessageArg{M: m}
	err = c.Cli.Call(ctx, "gregor.1.incoming.consumeMessage", []interface{}{__arg}, nil)
	return
}

func (c IncomingClient) ConsumePublishMessage(ctx context.Context, m Message) (err error) {
	__arg := ConsumePublishMessageArg{M: m}
	err = c.Cli.Call(ctx, "gregor.1.incoming.consumePublishMessage", []interface{}{__arg}, nil)
	return
}

func (c IncomingClient) Ping(ctx context.Context) (res string, err error) {
	err = c.Cli.Call(ctx, "gregor.1.incoming.ping", []interface{}{PingArg{}}, &res)
	return
}
