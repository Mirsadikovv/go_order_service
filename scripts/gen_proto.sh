#!/bin/bash

# Получаем текущий каталог из первого аргумента
CURRENT_DIR=$1

# Удаляем старые сгенерированные файлы
rm -rf "${CURRENT_DIR}/genproto/*"

# Устанавливаем пути к плагинам
PROTOC_GEN_GO=$(which protoc-gen-go)
PROTOC_GEN_GO_GRPC=$(which protoc-gen-go-grpc)

# Проверяем, что плагины найдены
for x in $(find ${CURRENT_DIR}/protos/* -type d); do
  protoc --plugin="protoc-gen-go=${PROTOC_GEN_GO}" \
         --plugin="protoc-gen-go-grpc=${PROTOC_GEN_GO_GRPC}" \
         -I=${x} -I=${CURRENT_DIR}/protos -I /usr/local/include \
         --go_out=${CURRENT_DIR} \
         --go-grpc_out=require_unimplemented_servers=false:${CURRENT_DIR} \
         ${x}/*.proto
done
