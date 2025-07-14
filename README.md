# challenge-stress-test

## Packages
- https://github.com/spf13/cobra

## Teste
- Criar a imagem:
~~~
docker build -t stress-test .
~~~
- Testando:
~~~
docker run --rm stress-test stress --url=https://httpbin.org/get --requests=50 --concurrency=5
~~~