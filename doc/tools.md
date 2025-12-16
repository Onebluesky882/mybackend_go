1. gen code pattern : go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

ใช้กับ Clean / DDD ยังไงดี?

โครงสร้างที่นิยมมาก 👇

internal/
api/ <- oapi-codegen output
domain/
usecase/
handler/ <- implement ServerInterface

Handler → call Usecase → Domain

9️⃣ ใช้กับ Makefile (แนะนำ)
generate:
oapi-codegen -generate types,chi-server -package api openapi.yaml > internal/api/api.gen.go

รัน:

make generate
a