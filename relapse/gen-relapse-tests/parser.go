//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"encoding/json"
	"encoding/xml"
	"github.com/gogo/protobuf/proto"
	descriptor "github.com/gogo/protobuf/protoc-gen-gogo/descriptor"
	"github.com/katydid/katydid/parser"
	jparser "github.com/katydid/katydid/parser/json"
	pparser "github.com/katydid/katydid/parser/proto"
	rparser "github.com/katydid/katydid/parser/reflect"
	xparser "github.com/katydid/katydid/parser/xml"
	"reflect"
)

type Codecs struct {
	Description string
	Parsers     map[string]NewParser
}

func getDesc(m interface{}) string {
	data, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ProtoEtc(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"reflect":   Reflect(m).Parsers["reflect"],
			"json":      Json(m).Parsers["json"],
			"protoName": ProtoName(m).Parsers["protoName"],
		},
	}
}

func Reflect(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"reflect": func() parser.Interface {
				return NewReflectParser(m)
			},
		},
	}
}

func Json(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"json": func() parser.Interface {
				return NewJsonParser(m)
			},
		},
	}
}

func JsonString(s string) Codecs {
	return Codecs{
		Description: s,
		Parsers: map[string]NewParser{
			"json": func() parser.Interface {
				return NewJsonStringParser(s)
			},
		},
	}
}

func ProtoName(m interface{}) Codecs {
	messageName := reflect.TypeOf(m).Elem().Name()
	packageName := "tests"
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"protoName": func() parser.Interface {
				return NewProtoNameParser(packageName, messageName, m.(ProtoMessage))
			},
		},
	}
}

func ProtoNum(m interface{}) Codecs {
	messageName := reflect.TypeOf(m).Elem().Name()
	packageName := "tests"
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"protoNum": func() parser.Interface {
				return NewProtoNumParser(packageName, messageName, m.(ProtoMessage))
			},
		},
	}
}

func XML(m interface{}) Codecs {
	return Codecs{
		Description: getDesc(m),
		Parsers: map[string]NewParser{
			"xml": func() parser.Interface {
				return NewXMLParser(m)
			},
		},
	}
}

func XMLString(s string) Codecs {
	return Codecs{
		Description: s,
		Parsers: map[string]NewParser{
			"xml": func() parser.Interface {
				return NewXMLStringParser(s)
			},
		},
	}
}

type NewParser func() parser.Interface

func NewReflectParser(m interface{}) parser.Interface {
	s := rparser.NewReflectParser()
	s.Init(reflect.ValueOf(m))
	return s
}

func NewJsonParser(m interface{}) parser.Interface {
	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := jparser.NewJsonParser()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	return s
}

func NewJsonStringParser(s string) parser.Interface {
	p := jparser.NewJsonParser()
	err := p.Init([]byte(s))
	if err != nil {
		panic(err)
	}
	return p
}

type ProtoMessage interface {
	Description() (desc *descriptor.FileDescriptorSet)
	proto.Message
}

func NewProtoNameParser(pkg, msg string, m ProtoMessage) parser.Interface {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := pparser.NewProtoNameParser(pkg, msg, m.Description())
	if err := s.Init(data); err != nil {
		panic(err)
	}
	return s
}

func NewProtoNumParser(pkg, msg string, m ProtoMessage) parser.Interface {
	data, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := pparser.NewProtoNumParser(pkg, msg, m.Description())
	if err := s.Init(data); err != nil {
		panic(err)
	}
	return s
}

func NewXMLParser(m interface{}) parser.Interface {
	data, err := xml.Marshal(m)
	if err != nil {
		panic(err)
	}
	s := xparser.NewXMLParser()
	err = s.Init(data)
	if err != nil {
		panic(err)
	}
	return s
}

func NewXMLStringParser(s string) parser.Interface {
	p := xparser.NewXMLParser()
	err := p.Init([]byte(s))
	if err != nil {
		panic(err)
	}
	return p
}
