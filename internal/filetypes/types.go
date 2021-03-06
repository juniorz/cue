// Code generated by gocode.Generate; DO NOT EDIT.

package filetypes

import (
	"fmt"

	"cuelang.org/go/cue"
	"cuelang.org/go/encoding/gocode/gocodec"
)

var cuegenvalFileInfo = cuegenMake("FileInfo", &FileInfo{})

// Validate validates x.
func (x *FileInfo) Validate() error {
	return cuegenCodec.Validate(cuegenvalFileInfo, x)
}

var cuegenCodec, cuegenInstance = func() (*gocodec.Codec, *cue.Instance) {
	var r *cue.Runtime
	r = &cue.Runtime{}
	instances, err := r.Unmarshal(cuegenInstanceData)
	if err != nil {
		panic(err)
	}
	if len(instances) != 1 {
		panic("expected encoding of exactly one instance")
	}
	return gocodec.New(r, nil), instances[0]
}()

// cuegenMake is called in the init phase to initialize CUE values for
// validation functions.
func cuegenMake(name string, x interface{}) cue.Value {
	f, err := cuegenInstance.LookupField(name)
	if err != nil {
		panic(fmt.Errorf("could not find type %q in instance", name))
	}
	v := f.Value
	if x != nil {
		w, err := cuegenCodec.ExtractType(x)
		if err != nil {
			panic(err)
		}
		v = v.Unify(w)
	}
	return v
}

// Data size: 1051 bytes.
var cuegenInstanceData = []byte("\x01\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcWKo\u0736\x13\x17\xd7\xfe\x03\u007f\x12q\xbfA\x01U\x87\xc2\xd9\xc3^\v,\x10\xf8\x92\x06\u0225(z5\x02\x83\x96(Y\x8d\x96\x14$.`#\xbb\x87\xb6i\xdaK\xbfr\xb6\x18\x0e\x1f\"\xa55\x9a\x1e\xca\xd3\xea7\x9c7\xe7\xb1W\xa7?Vdu\xfa3#\xa7_\xb3\xec\xbb\xd3/\x17\x84\xbch\u5a39,\xc5k\xae9\xe0\xe4\x82\\\xfe\xa4\x94&\xab\x8c\\\xfe\xc8\xf5\x03y\x91\x91\xff\xbdi;1\x92\u04e7,\u02fe>\xfd\xbe\"\xe4\xab\xdbw\xe5^l\uadb3\x9c\x9f2r\xfa\x98e\u05e7\xdf.\b\xf9\u007f\xc0?fdE.\u007f\xe0;\x01\x82.\r\u0232,\xfb|\xf5\x17XB\u020a\x10\xaa\x9fz1n\u02bd \x9f\xaft\xcf\xcb\xf7\xbc\x11\xf9\xfd\xbe\xed*\xc6@u\xbe\xdd\xe6\x1f\x18\x05\xa9\x92\xef\xc46\xb7g\xd4C+\x1bF\x85,U\xd5\xca\xc6\x13\xbe\xb7\x00\xa3\xad\xd4b\xe8\a\xa1\xb9n\x95\xbc\xd9\xe6o#\x80\xd1Z\r\xbb\x1b\u03d8\xe7\xf9\x1b5\xec\x18\u057c\x19o\x8cVz\x8bj\xdem\xbd\xbe#;2\xa7\x02l\xc3\xf3\u036b\xa2`\xb1x Z&\x10\x1b\xeeN\xac\aMF\x91\x16\x8f\x1a5\x06\u007f\n\x00\vF\x8d\x99\xc8\\T\\\xf3\x02\x8c\xa0\xf0\v9\x90<!\x95{1\x93U\xee\x05\x12\xc7\xf2A\xecbN\x84\x90\xfc\xf3\xa8\xe4\x8c\x19@ \xe7\xdf\xe6\xd7\xeb\xc0\xb8.\x8a\xfc\x10\x14\xe7\a\xc3\x17G\x1dn\x01\xbb\u0551\x1f\xf2\xbb\xc8#<\xeb\"\u040b\xba\x95\xbc\x03\x81/\u045cn\u045e\x0e\xed}\xe2\xbb9\x1d@$7*%\xe2)JUA@f\xd6\x16\x00\xfa\xb4P\xdaq#\xb1Q\x80\x1f\x8d\xcc~Pz&\xb60(*5\x0eD\x11\xf6.1\xda\f\xbc\u007f\x88\x88\x06qIm\x92\x9c66\xa5\xaa\x9a\xe5\xf4_y\u27b1\xf5%\xa4f\xbb\x98\xbci\xee\f\x83\xea\x85\xe4}{\u6da5\x16X&P\xbeoe\xadl\t\xe3\x8b\xf5)\xd7\xc3^\u41fc\xe6\xdd(\x18\x1dD-\x06!K1n\xe7\xc4\xf2\xa9\uc430\xc0Y\x89\xba\x95-\x18\x007\xee\x95\xea\xc0J\xf8\x86$\xc0A\xacTr\xd4\x03o\xa5\x0e\xf7\xde\v\u047f\x165\xdfw\x00\"\xd6\xcaR\xed\xfaNh\xd3k,\xb6\xeb\u0560\x9d\x05\x88\x8dz\x10\u073fb\xc4*U\x8e\xc1E\u0138\xd6C{\xbf\xd7\u8035}m\x8d\x87\x10\xb1#\u06e9J`\x9eZ\xd9\xefm'H\xc2G\xad\xf4\xc4}J7\x9b\r\xe6\u04ff\r\x97r\xdf\b(]\u00cb\x1a7\x98J[\x85\x86\u01fc\x03\xa8\xec\x0f\x8c\x8aGp\xf3Y\xed\x89\xfd_\xa6\x1d^\xc09\u0755\xa8\xff{\xb7Y\xcc\xea\x19]n\xad2od\xe8\x8e\xee\x06\x06\"?\x98\xc00\xea_\x80338>\x11\x13\x9aV\xa4\b\xe2j\xc5 \x9b\xe9\f\x89\xfa\x841\xba\x9f\xa8\tC%\xf1\xe7\xccu\xb5h\u0579\xeb\u03f7A:\r\xfc\x84+4\xb2e%\x13\x06f\x10s\xf9\x16\x96\x88W~\x16\xbb'\xe2\x84\x16\xc56\xbfs\x1f\xf3\x01\xf7\u0314\x81_\xae\xff&\x9d$}pqO\xb9\x8e\xc8/\xcdd\x8c\x10F\x93\x8e\x93\u028b{OJ\x8d\xbb\u040c\x1a\xf5\xa3\x94\x1aw\xa6\xa4b}&fS\xca\x1d?\xad\x12\x97\x97\r?\x9b>\xd42\x1bw!\x19\x18w\xdc \x9a\xe9&\x91L\nWYQv\x96\xb3\xe2\xd08\xf2\xcf\x1b\x1eGz9\xc2i\xec\x929m\xf6?\xff\x8c\u0734\x8a#\x93\xd6r\xba\xbf\x85E\x12\x87~<\x11\xbdc\xd3I\xb8\x18\x83\xc5\x10,z\x95V\xb7]n\u0163\x16r\xc40\xdb\u06b2\xe7\x8e\xd1\x026uD`\xb5\x80/\x00\u0342\xb8u |9\xb4\x03\u0623\x1d\xc0]e\xaf\u01f0\\\x86\xcd.\xe7E\u00d7A-\x18\xa3\xfaQOPh\u007f\x806\u02ba`\xd0F\x01\x86m\u0299f\xbe\xa0\u1d1d\b\xdd\u0155\x89\xfb\xe3Q\xd4Jm\xec*<Y\xc9\xddvtd\xf1:\xf4\xe5\x9d+]\xc6B[]\xe3J;\u065e\x17J-Z\xcd\xfe1/\u02f2\xbf\x03\x00\x00\xff\xff\xea*w\xba\x1d\x0e\x00\x00")
