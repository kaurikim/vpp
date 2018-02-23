// Code generated by github.com/ungerik/pkgreflect DO NOT EDIT.

package dhcp

import "reflect"

var Types = map[string]reflect.Type{
	"DhcpClientConfig":      reflect.TypeOf((*DhcpClientConfig)(nil)).Elem(),
	"DhcpClientConfigReply": reflect.TypeOf((*DhcpClientConfigReply)(nil)).Elem(),
	"DhcpComplEvent":        reflect.TypeOf((*DhcpComplEvent)(nil)).Elem(),
	"DhcpProxyConfig":       reflect.TypeOf((*DhcpProxyConfig)(nil)).Elem(),
	"DhcpProxyConfigReply":  reflect.TypeOf((*DhcpProxyConfigReply)(nil)).Elem(),
	"DhcpProxyDetails":      reflect.TypeOf((*DhcpProxyDetails)(nil)).Elem(),
	"DhcpProxyDump":         reflect.TypeOf((*DhcpProxyDump)(nil)).Elem(),
	"DhcpProxySetVss":       reflect.TypeOf((*DhcpProxySetVss)(nil)).Elem(),
	"DhcpProxySetVssReply":  reflect.TypeOf((*DhcpProxySetVssReply)(nil)).Elem(),
	"DhcpServer":            reflect.TypeOf((*DhcpServer)(nil)).Elem(),
}

var Functions = map[string]reflect.Value{
	"NewDhcpClientConfig":      reflect.ValueOf(NewDhcpClientConfig),
	"NewDhcpClientConfigReply": reflect.ValueOf(NewDhcpClientConfigReply),
	"NewDhcpComplEvent":        reflect.ValueOf(NewDhcpComplEvent),
	"NewDhcpProxyConfig":       reflect.ValueOf(NewDhcpProxyConfig),
	"NewDhcpProxyConfigReply":  reflect.ValueOf(NewDhcpProxyConfigReply),
	"NewDhcpProxyDetails":      reflect.ValueOf(NewDhcpProxyDetails),
	"NewDhcpProxyDump":         reflect.ValueOf(NewDhcpProxyDump),
	"NewDhcpProxySetVss":       reflect.ValueOf(NewDhcpProxySetVss),
	"NewDhcpProxySetVssReply":  reflect.ValueOf(NewDhcpProxySetVssReply),
}

var Variables = map[string]reflect.Value{}

var Consts = map[string]reflect.Value{
	"VlAPIVersion": reflect.ValueOf(VlAPIVersion),
}
