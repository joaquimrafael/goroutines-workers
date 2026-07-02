# goroutines-workers

Pool de workers em Go: processa uma fila de tarefas usando um número fixo de goroutines que se comunicam por channels.

O lema do Go na prática: *"não comunique compartilhando memória; compartilhe memória comunicando."*

## Como funciona

- Um channel `jobs` (entrada) e um channel `results` (saída).
- `N` workers rodam como goroutines, cada um consumindo `jobs` com `for n := range jobs`.
- Uma goroutine produtora envia todos os jobs e dá `close(jobs)` — o sinal de "acabou".
- Um `WaitGroup` + uma goroutine coordenadora fazem `wg.Wait(); close(results)` na hora certa.
- O consumidor lê `results` com `for r := range results` e encerra sozinho quando o channel fecha.

`RunPool(in []int, n int) []int` recebe a entrada e o número de workers e devolve os resultados (aqui, o quadrado de cada número). A ordem dos resultados varia — é o comportamento esperado em concorrência.

## Rodar

```sh
go run .
```

## Testar

```sh
go test -race -timeout 5s
```

- `-race` pega acesso concorrente errado.
- `-timeout 5s` pega deadlock (falha em vez de travar pra sempre).
