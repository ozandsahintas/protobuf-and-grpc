package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	any2 "protocol-buffers/proto/code-gen/any"
	px "protocol-buffers/proto/code-gen/complex"
	pe "protocol-buffers/proto/code-gen/enum"
	"protocol-buffers/proto/code-gen/file"
	pm "protocol-buffers/proto/code-gen/maps"
	pof "protocol-buffers/proto/code-gen/oneof"
	ps "protocol-buffers/proto/code-gen/simple"
	ttt "protocol-buffers/proto/code-gen/test-dir"
	"reflect"
)

func main() {

	//fmt.Println(doSimple())
	//fmt.Println(doComplex())
	//fmt.Println(doEnum())

	//doOneof(&pof.Result_Id{Id: 20})
	//doOneof(&pof.Result_Msg{Msg: "A message"})
	//doOneof("hey")

	//fmt.Println(doMap())
	//fmt.Println(doMap().Ids["firstId"].Id)

	//doFile(doSimple())

	//test()

	doAny()

	//js := doToJSON(doSimple())
	//m := doFromJSON(js, reflect.TypeOf(ps.Simple{}))
	//fmt.Println(js)
	//fmt.Println(m)

}

func doSimple() *ps.Simple {
	return &ps.Simple{
		Number:     12,
		IsSimple:   true,
		Name:       "Simple",
		SampleList: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *px.Complex {
	return &px.Complex{
		Other: &px.Other{Id: -2, Name: "One Other"},
		Others: []*px.Other{
			{Id: -3, Name: "Second Other"},
			{Id: -4, Name: "Third Other"},
		},
	}
}

func doEnum() *pe.Enum {
	return &pe.Enum{
		//EyeColor:       pe.EyeColor_EYE_COLOR_UNKNOWN, -> Not printed
		EyeColor:       pe.EyeColor_EYE_COLOR_BROWN,
		EyeColorNumber: 1,
	}
}

func doOneof(message interface{}) {
	switch x := message.(type) {
	case *pof.Result_Id:
		fmt.Printf("Int: %d\n", message.(*pof.Result_Id).Id)
	case *pof.Result_Msg:
		fmt.Printf("String: %s\n", message.(*pof.Result_Msg).Msg)
	default:
		fmt.Printf("unexpected: %v", x)
	}
}

func doMap() *pm.AMap {
	return &pm.AMap{Ids: map[string]*pm.IdWrapper{
		"firstId":  {Id: 1},
		"secondId": {Id: 2},
		"thirdId":  {Id: 3},
		"fourthId": {Id: 4},
		"fifthId":  {Id: 5},
	}}
}

func test() {

	address := ttt.Address{
		City: &ttt.City{
			Name:       "City Name",
			ZipCode:    "12345",
			CountyName: "County Name",
		},
		Street: &ttt.City_Street{
			Name: "Street Name",
			City: &ttt.City{
				Name:       "",
				ZipCode:    "",
				CountyName: "",
			},
		},
		Building: &ttt.City_Street_Building{
			Name:   "Building Name",
			Number: 123,
			Street: &ttt.City_Street{
				Name: "Street Name",
				City: &ttt.City{
					Name:       "City Name",
					ZipCode:    "12345",
					CountyName: "County Name",
				},
			},
		},
	}

	fmt.Println(doToJSON(&address))
}

func doFile(p proto.Message) {
	path := "proto/code-gen/file/simple.bin"

	file.WriteFile(path, p)
	msg := &ps.Simple{}
	file.ReadFile(path, msg)
	fmt.Println(msg)
}

func doAny() {
	var a anypb.Any

	anypb.MarshalFrom(&a, &any2.StructOne{
		Id:     "Id",
		Number: 1,
	}, proto.MarshalOptions{})

	fmt.Println(&a)

	s3 := &any2.StructOne{}
	if err := a.UnmarshalTo(s3); err != nil {
		fmt.Println(err)
	}

	fmt.Println(s3)

	anypb.MarshalFrom(&a, &any2.StructTwo{
		Id:      "NewId",
		AString: "A String",
	}, proto.MarshalOptions{})

	fmt.Println(&a)

}

func doToJSON(p proto.Message) string {
	jsonString := file.ToJSON(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	msg := reflect.New(t).Interface().(proto.Message)
	file.FromJSON(jsonString, msg)
	return msg
}
