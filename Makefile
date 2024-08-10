# Wrote by yijian on 2024/08/09

TARGET=gadget_crafted_service
CLIENT=gadget_crafted_client

ALL: ${TARGET} ${CLIENT}

${TARGET}: gadgetcrafted.go
	go build -o $@ $<

${CLIENT}: gadget_crafted_client.go
	go build -o $@ $<

api: gadget_crafted.api
	goctl api go --style=go_zero -api gadget_crafted.api -dir .
	go mod tidy

tidy:
	go mod tidy

clean:
	rm -f ${TARGET} ${CLIENT}
