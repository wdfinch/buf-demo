#!/bin/bash
input=$(echo '{"dog_id": '"$1"'}')
output=$(grpcurl \
	-protoset <(cd ../proto && buf build -o -) \
	-plaintext \
	-d "$input" \
	localhost:50051 petland.dog.v1.DogApiService/GetDog)

echo "$output" | jj -p
