generate:
	protoc --go_opt=module=${package} --go_out=. proto/code-gen/**/*.proto

decode-raw:
	cat proto/protoc-adv/simple.pb | protoc --decode_raw


decode-file:
	cat proto/protoc-adv/simple.pb | protoc --decode=Simple proto/protoc-adv/simple.proto > proto/protoc-adv/simple.txt

decode:
	cat proto/protoc-adv/simple.pb | protoc --decode=Simple proto/protoc-adv/simple.proto


encode:
	cat proto/protoc-adv/simple.txt | protoc --encode=Simple proto/protoc-adv/simple.proto > proto/protoc-adv/simple.pb

