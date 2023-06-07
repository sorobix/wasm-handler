# WASM-Handler

Go App that listens to websocket connections at:
- /format: to format the Go code
- /compile: to compile the input code into a WASM (throws error if it fails)

## Format Rust code
Message Format:
- Input Base64(input rust code)
- Output Base64({Success:true/false, Data:(Base64 encoded input rust formatted code)})
```
Input: IyFbbm9fc3RkXQp1c2Ugc29yb2Jhbl9zZGs6Ontjb250cmFjdGltcGwsIHZlYywgRW52LCBTeW1ib2wsIFZlY307CnB1YiBzdHJ1Y3QgQ29udHJhY3Q7CiNbY29udHJhY3RpbXBsXQppbXBsIENvbnRyYWN0IHsKICAgICAgICAgICAgICBwdWIgZm4gaGVsbG8oZW52OiBFbnYsIHRvOiBTeW1ib2wpIC0+IFZlYzxTeW1ib2w+IHsKICAgICAgICB2ZWMhWyZlbnYsIFN5bWJvbDo6c2hvcnQoIkhlbGxvIiksIHRvXQogICAgfQp9Cg==

Output:
eyJTdWNjZXNzIjp0cnVlLCJEYXRhIjoiSXlGYmJtOWZjM1JrWFFwMWMyVWdjMjl5YjJKaGJsOXpaR3M2T250amIyNTBjbUZqZEdsdGNHd3NJSFpsWXl3Z1JXNTJMQ0JUZVcxaWIyd3NJRlpsWTMwN0NuQjFZaUJ6ZEhKMVkzUWdRMjl1ZEhKaFkzUTdDaU5iWTI5dWRISmhZM1JwYlhCc1hRcHBiWEJzSUVOdmJuUnlZV04wSUhzS0lDQWdJSEIxWWlCbWJpQm9aV3hzYnlobGJuWTZJRVZ1ZGl3Z2RHODZJRk41YldKdmJDa2dMVDRnVm1WalBGTjViV0p2YkQ0Z2V3b2dJQ0FnSUNBZ0lIWmxZeUZiSm1WdWRpd2dVM2x0WW05c09qcHphRzl5ZENnaVNHVnNiRzhpS1N3Z2RHOWRDaUFnSUNCOUNuMEsifQ
```

## Compile WASM from Rust code
Message Format:
- Input {CargoToml:Base64(cargoToml), MainRs:Base64(rust_file)}
- Output Base64({Id: Redis_key, Success:true/false, Message:Compilation Output})

```
Input:{"cargoToml":"W3BhY2thZ2VdCm5hbWUgPSAic29yb2JpeF90ZW1wIgp2ZXJzaW9uID0gIjAuMS4wIgplZGl0aW9uID0gIjIwMjEiCgpbbGliXQpjcmF0ZS10eXBlID0gWyJjZHlsaWIiXQoKW2ZlYXR1cmVzXQp0ZXN0dXRpbHMgPSBbInNvcm9iYW4tc2RrL3Rlc3R1dGlscyJdCgpbZGVwZW5kZW5jaWVzXQpzb3JvYmFuLXNkayA9ICIwLjguNCIKCltkZXZfZGVwZW5kZW5jaWVzXQpzb3JvYmFuLXNkayA9IHsgdmVyc2lvbiA9ICIwLjguNCIsIGZlYXR1cmVzID0gWyJ0ZXN0dXRpbHMiXSB9CgpbcHJvZmlsZS5yZWxlYXNlXQpvcHQtbGV2ZWwgPSAieiIKb3ZlcmZsb3ctY2hlY2tzID0gdHJ1ZQpkZWJ1ZyA9IDAKc3RyaXAgPSAic3ltYm9scyIKZGVidWctYXNzZXJ0aW9ucyA9IGZhbHNlCnBhbmljID0gImFib3J0Igpjb2RlZ2VuLXVuaXRzID0gMQpsdG8gPSB0cnVlCgpbcHJvZmlsZS5yZWxlYXNlLXdpdGgtbG9nc10KaW5oZXJpdHMgPSAicmVsZWFzZSIKZGVidWctYXNzZXJ0aW9ucyA9IHRydWU=", "MainRs":"IyFbbm9fc3RkXQp1c2Ugc29yb2Jhbl9zZGs6Ontjb250cmFjdGltcGwsIHZlYywgRW52LCBTeW1ib2wsIFZlY307CgpwdWIgc3RydWN0IENvbnRyYWN0OwoKI1tjb250cmFjdGltcGxdCmltcGwgQ29udHJhY3QgewogICAgcHViIGZuIGhlbGxvKGVudjogRW52LCB0bzogU3ltYm9sKSAtPiBWZWM8U3ltYm9sPiB7CiAgICAgICAgdmVjIVsmZW52LCBTeW1ib2w6OnNob3J0KCJIZWxsbyIpLCB0b10KICAgIH0KfQ=="}

Output: 
eyJJZCI6IjNiYzJkM2VhMDhjOGY4N2M1NWI2MWUzODUwNjczZDcwIiwiU3VjY2VzcyI6dHJ1ZSwiTWVzc2FnZSI6IkNvbXBpbGF0aW9uIHN1Y2Nlc3NmdWwhIn0=
```

## Redis

It also uses Redis as a caching service for the compile functionality to prevent same repeatative builds also leading to faster response times.

Redis Data Format:
```
key: md5(base64(cargoToml+mainRs))

value:
    base64
    (
        Success: true/false,
        Message: base64(compiler output),
        Wasm: base64(wasm file content)
    )
```
