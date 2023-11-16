# Comparação de Desempenho Serialização e Desserialização do JSON no Fiber

A ideia desse projeto é comparar o desempenho de uma API desenvolvida com o web-framework Fiber.
Nessa comparação temos duas APIs que utilizam o mesmo handler para receber uma requisicação POST e precisa fazer a decodificação da Request recebida no body e realizar a 
codificação de um Response.
A diferença entre as duas APIs fica por conta da configuração do fiber, em uma das implementações foi utilizada a configuração padrão e na outra foi definido que a biblioteca `json`
utilizada será a [goccy/go-json](https://github.com/goccy/go-json).
- Exemplo:
```go
    import (
        "github.com/goccy/go-json"
    )

    app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,
	})
```
## Como Testar?
Para fazer essa comparação foi utilizado o **[Bombardier](https://github.com/codesenberg/bombardier)**, uma ferramenta de benchmarking HTTP desenvolvida em Go.

O teste foi realizado utilizando os seguintes cenários:

Cenário 1:
- Quantidade de Usuários Virtuais: 100
- Quantidade de Requisições: 10000000

Cenário 2:
- Quantidade de Usuários Virtuais: 100
- Tempo de Execução: 60 segundos

Para rodar o teste é necessário subir as aplicações e executar o seguinte comando:
- Cenário 1

```bombardier -c 100 -n 10000000 -m POST  -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" -b req.json http://localhost:3000```

- Cenário 2

```bombardier -c 100 -d 60s -m POST  -H "Content-Type: application/x-www-form-urlencoded; charset=UTF-8" -b req.json http://localhost:3000```

## Resultados:

### Biblioteca Padrão do Go ([encoding/json](https://pkg.go.dev/encoding/json))
#### Cenário 1:

```
Bombarding http://localhost:3000 with 10000000 request(s) using 100 connection(s)
 10000000 / 10000000 [=====================================================] 100.00% 158287/s 1m3s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    158887.12   27938.18  233398.71
  Latency      625.94us   500.97us    85.38ms
  HTTP codes:
    1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    43.76MB/s
```

#### Cenário 2:

```
Bombarding http://localhost:3000 for 1m0s using 100 connection(s)
[==========================================================================] 1m0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    162881.09   22906.63  232135.89
  Latency      609.31us   311.12us    98.46ms
  HTTP codes:
    1xx - 0, 2xx - 9785295, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    42.93MB/s
```

### Biblioteca go-json ([goccy/go-json](https://github.com/goccy/go-json))

#### Cenário 1:

```
Bombarding http://localhost:3000 with 10000000 request(s) using 100 connection(s)
 10000000 / 10000000 [=====================================================] 100.00% 162981/s 1m1s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    163121.35   28614.96  230859.70
  Latency      608.68us   533.98us   120.71ms
  HTTP codes:
    1xx - 0, 2xx - 10000000, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    44.99MB/s
```

#### Cenário 2:
 ```
 Bombarding http://localhost:3000 for 1m0s using 100 connection(s)
[==========================================================================] 1m0s
Done!
Statistics        Avg      Stdev        Max
  Reqs/sec    166746.05   25303.72  237240.97
  Latency      595.35us   348.96us    77.93ms
  HTTP codes:
    1xx - 0, 2xx - 10014291, 3xx - 0, 4xx - 0, 5xx - 0
    others - 0
  Throughput:    43.93MB/s
 ```

## Conclusão:

### Cenário 1

| Biblioteca  | Reqs/sec (média) | Throughput |  Latência (média) |  Tempo |
| ------------- | ------------- | ------------- | ------------- | ------------- |
|  encoding/json  | 158887.12  | 43.76MB/s  | 625.94us  | 1m3s  |
| goccy/go-json  | 163121.35  | 44.99MB/s  | 608.68us  | 1m1s  |


### Cenário 2

| Biblioteca  | Reqs/sec (média) | Throughput |  Latência (média) |  Total de Requisições |
| ------------- | ------------- | ------------- | ------------- | ------------- |
|  encoding/json  | 162881.09  | 42.93MB/s  | 609.31us  | 9785295  |
| goccy/go-json  | 166746.05  | 43.93MB/s  | 595.35us  | 10014291  |

