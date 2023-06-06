# WASM-Handler

Go App that listens to websocket connections at:
- /format: to format the Go code
- /compile: to compile the input code into a WASM (throws error if it fails)

## Format Rust code
Message Format:
- Input "(Base64 encoded input rust code)"
- Output {Success:true/false, Data: "(Base64 encoded input rust formatted code)"}
```
Input: IyFbbm9fc3RkXQp1c2Ugc29yb2Jhbl9zZGs6Ontjb250cmFjdGltcGwsIHZlYywgRW52LCBTeW1ib2wsIFZlY307CnB1YiBzdHJ1Y3QgQ29udHJhY3Q7CiNbY29udHJhY3RpbXBsXQppbXBsIENvbnRyYWN0IHsKICAgICAgICAgICAgICBwdWIgZm4gaGVsbG8oZW52OiBFbnYsIHRvOiBTeW1ib2wpIC0+IFZlYzxTeW1ib2w+IHsKICAgICAgICB2ZWMhWyZlbnYsIFN5bWJvbDo6c2hvcnQoIkhlbGxvIiksIHRvXQogICAgfQp9Cg==

Output:
{"Success":true,"Data":"IyFbbm9fc3RkXQp1c2Ugc29yb2Jhbl9zZGs6Ontjb250cmFjdGltcGwsIHZlYywgRW52LCBTeW1ib2wsIFZlY307CnB1YiBzdHJ1Y3QgQ29udHJhY3Q7CiNbY29udHJhY3RpbXBsXQppbXBsIENvbnRyYWN0IHsKICAgIHB1YiBmbiBoZWxsbyhlbnY6IEVudiwgdG86IFN5bWJvbCkgLT4gVmVjPFN5bWJvbD4gewogICAgICAgIHZlYyFbJmVudiwgU3ltYm9sOjpzaG9ydCgiSGVsbG8iKSwgdG9dCiAgICB9Cn0K"}
```

## Compile WASM from Rust code
Message Format:
- Input {"CargoToml":"cargo toml file", "MainRs":"rust file"}
- Output {"Id": "Redis key", "Success":true/false, "Message":"Compilation Output"}

```
Input: {"CargoToml":"[package]\nname = \"sorobix_temp\"\nversion = \"0.1.0\"\nedition = \"2021\"\n[lib]\ncrate-type = [\"cdylib\"]\n[dependencies]\nsoroban-sdk = { version = \"0.7.0\"}\n[profile.release]\nopt-level = \"z\"\noverflow-checks = true\ndebug = 0\nstrip = \"symbols\"\ndebug-assertions = false\npanic = \"abort\"\ncodegen-units = 1\nlto = true\n[profile.release-with-logs]\ninherits = \"release\"\ndebug-assertions = true","MainRs":"\n#![no_std]\nuse soroban_sdk::{contractimpl, vec, Env, Symbol, Vec};\npub struct Contract;\n#[contractimpl]\nimpl Contract {\n\tpub fn wello(env: Env, tos: Symbol) -> Vec<Symbol> {\n\t\tvec![&env, Symbol::(\"Hello\"), to]\n\t}\n}"}

Output: 
{"Id":"a8034469781ef9d7a01c54c98b3c6b0b","Success":false,"Message":"failed to compile rust project: Updating crates.io index\n   Compiling proc-macro2 v1.0.59\n   Compiling quote v1.0.28\n   Compiling unicode-ident v1.0.9\n   Compiling serde v1.0.163\n   Compiling serde_json v1.0.96\n   Compiling ryu v1.0.13\n   Compiling itoa v1.0.6\n   Compiling autocfg v1.1.0\n   Compiling fnv v1.0.7\n   Compiling strsim v0.10.0\n   Compiling ident_case v1.0.1\n   Compiling syn v1.0.109\n   Compiling typenum v1.16.0\n   Compiling version_check v0.9.4\n   Compiling thiserror v1.0.40\n   Compiling either v1.8.1\n   Compiling indexmap v1.9.3\n   Compiling itertools v0.10.5\n   Compiling num-traits v0.2.15\n   Compiling num-integer v0.1.45\n   Compiling hashbrown v0.12.3\n   Compiling generic-array v0.14.7\n   Compiling prettyplease v0.1.25\n   Compiling num-bigint v0.4.3\n   Compiling cpufeatures v0.2.7\n   Compiling cfg-if v1.0.0\n   Compiling ethnum v1.3.2\n   Compiling syn v2.0.18\n   Compiling static_assertions v1.1.0\n   Compiling wasmparser v0.88.0\n   Compiling base64 v0.13.1\n   Compiling crypto-common v0.1.6\n   Compiling block-buffer v0.10.4\n   Compiling digest v0.10.7\n   Compiling sha2 v0.10.6\n   Compiling darling_core v0.20.1\n   Compiling darling_core v0.14.4\n   Compiling serde_derive v1.0.163\n   Compiling thiserror-impl v1.0.40\n   Compiling bytes-lit v0.0.4\n   Compiling darling_macro v0.20.1\n   Compiling darling_macro v0.14.4\n   Compiling darling v0.20.1\n   Compiling darling v0.14.4\n   Compiling serde_with_macros v2.3.3\n   Compiling hex v0.4.3\n   Compiling serde_with v2.3.3\n   Compiling crate-git-revision v0.0.4\n   Compiling stellar-xdr v0.0.15\n   Compiling soroban-env-common v0.0.15\n   Compiling soroban-spec v0.7.0\n   Compiling soroban-env-macros v0.0.15\n   Compiling soroban-env-guest v0.0.15\n   Compiling soroban-sdk-macros v0.7.0\n   Compiling soroban-sdk v0.7.0\n   Compiling sorobix_temp v0.1.0 (/tmp/rust-project2911637337/sorobix_temp)\nerror: expected type, found `\"Hello\"`\n  --\u003e src/lib.rs:8:23\n   |\n8  |         vec![\u0026env, Symbol::(\"Hello\"), to]\n   |                             ^^^^^^^ expected type\n   |\n  ::: /home/shubham/.cargo/registry/src/index.crates.io-6f17d22bba15001f/soroban-sdk-0.7.0/src/vec.rs:43:19\n   |\n43 |     ($env:expr, $($x:expr),+ $(,)?) =\u003e {\n   |                   ------- while parsing argument for this `expr` macro fragment\n\nerror: could not compile `sorobix_temp` (lib) due to previous error"}
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
