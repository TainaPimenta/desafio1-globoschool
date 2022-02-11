# desafio1-globoschool

<h1 align="center">  

![spotplay](https://media0.giphy.com/media/xT5LMVDzXfBcaPIXNC/giphy.gif?cid=790b76116515cdad633d8692d057fb3e8463edd6d3503d00&rid=giphy.gif&ct=g)
  
</h1>

<h1>O Projeto</h1>
#Primeiro desafio criando GO, VScode, docker e git;

#O projeto é feito em Go, que é uma linguagem de programação de código aberto.

 <h4 align="center" > 
 <img src="https://cdn.dribbble.com/users/1792477/screenshots/6816387/ezgif.com-resize__3__still_2x.gif?compress=1&resize=400x300" width="180">
 
  
 <h1>Como executar</h1> 
Podemos executá-lo de maneira bem simples:

#No seu terminal digite o comando go run main.go;

#Aparecerá uma mensagem para que você digite o seu nome e depois de digitado o seu nome irá junto a uma mensagem;

#Digite algum nome e verá ele ser execultado em seu terminal.


<h1>Dockerfile</h1>
Você também poderá executar através do dockerfile:

#Criando a imagem
-docker build -t nome da imagem -f ./Docker/golang/Dockerfile . 

#Bildando a imagem
-docker run --name nome da imagem --rm -v $PWD:/app:Z -w /app projeto-go go build -o ./bin/cmd ./main.go 

#Criando um binário
-docker run -it --name projeto-go-cont --rm -v $PWD:/app:Z -w /app nome da imagem ./bin/cmd 


