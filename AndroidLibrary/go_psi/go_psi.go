// Package go_psi is an autogenerated binder stub for package psi.
//   gobind -lang=go github.com/Psiphon-Labs/psiphon-tunnel-core/AndroidLibrary/psi
//
// File is generated by gobind. Do not edit.
package go_psi

import (
	"github.com/Psiphon-Labs/psiphon-tunnel-core/AndroidLibrary/psi"
	"golang.org/x/mobile/bind/seq"
)

const (
	proxyPsiphonProviderDescriptor                 = "go.psi.PsiphonProvider"
	proxyPsiphonProviderBindToDeviceCode           = 0x10a
	proxyPsiphonProviderGetDnsServerCode           = 0x20a
	proxyPsiphonProviderHasNetworkConnectivityCode = 0x30a
	proxyPsiphonProviderNoticeCode                 = 0x40a
)

func proxyPsiphonProviderBindToDevice(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(psi.PsiphonProvider)
	param_fileDescriptor := in.ReadInt()
	err := v.BindToDevice(param_fileDescriptor)
	if err == nil {
		out.WriteUTF16("")
	} else {
		out.WriteUTF16(err.Error())
	}
}

func proxyPsiphonProviderGetDnsServer(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(psi.PsiphonProvider)
	res := v.GetDnsServer()
	out.WriteUTF16(res)
}

func proxyPsiphonProviderHasNetworkConnectivity(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(psi.PsiphonProvider)
	res := v.HasNetworkConnectivity()
	out.WriteInt(res)
}

func proxyPsiphonProviderNotice(out, in *seq.Buffer) {
	ref := in.ReadRef()
	v := ref.Get().(psi.PsiphonProvider)
	param_noticeJSON := in.ReadUTF16()
	v.Notice(param_noticeJSON)
}

func init() {
	seq.Register(proxyPsiphonProviderDescriptor, proxyPsiphonProviderBindToDeviceCode, proxyPsiphonProviderBindToDevice)
	seq.Register(proxyPsiphonProviderDescriptor, proxyPsiphonProviderGetDnsServerCode, proxyPsiphonProviderGetDnsServer)
	seq.Register(proxyPsiphonProviderDescriptor, proxyPsiphonProviderHasNetworkConnectivityCode, proxyPsiphonProviderHasNetworkConnectivity)
	seq.Register(proxyPsiphonProviderDescriptor, proxyPsiphonProviderNoticeCode, proxyPsiphonProviderNotice)
}

type proxyPsiphonProvider seq.Ref

func (p *proxyPsiphonProvider) BindToDevice(fileDescriptor int) error {
	in := new(seq.Buffer)
	in.WriteInt(fileDescriptor)
	out := seq.Transact((*seq.Ref)(p), proxyPsiphonProviderBindToDeviceCode, in)
	res_0 := out.ReadError()
	return res_0
}

func (p *proxyPsiphonProvider) GetDnsServer() string {
	in := new(seq.Buffer)
	out := seq.Transact((*seq.Ref)(p), proxyPsiphonProviderGetDnsServerCode, in)
	res_0 := out.ReadUTF16()
	return res_0
}

func (p *proxyPsiphonProvider) HasNetworkConnectivity() int {
	in := new(seq.Buffer)
	out := seq.Transact((*seq.Ref)(p), proxyPsiphonProviderHasNetworkConnectivityCode, in)
	res_0 := out.ReadInt()
	return res_0
}

func (p *proxyPsiphonProvider) Notice(noticeJSON string) {
	in := new(seq.Buffer)
	in.WriteUTF16(noticeJSON)
	seq.Transact((*seq.Ref)(p), proxyPsiphonProviderNoticeCode, in)
}

func proxy_Start(out, in *seq.Buffer) {
	param_configJson := in.ReadUTF16()
	param_embeddedServerEntryList := in.ReadUTF16()
	var param_provider psi.PsiphonProvider
	param_provider_ref := in.ReadRef()
	if param_provider_ref.Num < 0 { // go object
		param_provider = param_provider_ref.Get().(psi.PsiphonProvider)
	} else { // foreign object
		param_provider = (*proxyPsiphonProvider)(param_provider_ref)
	}
	err := psi.Start(param_configJson, param_embeddedServerEntryList, param_provider)
	if err == nil {
		out.WriteUTF16("")
	} else {
		out.WriteUTF16(err.Error())
	}
}

func proxy_Stop(out, in *seq.Buffer) {
	psi.Stop()
}

func init() {
	seq.Register("psi", 1, proxy_Start)
	seq.Register("psi", 2, proxy_Stop)
}