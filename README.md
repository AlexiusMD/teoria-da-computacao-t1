# Máquina de Turing Universal

Este repositório é referente à implementação de uma *Máquina de Turing Universal (MTU)*, para a disciplina de Teoria da Computabilidade e Complexidade, ministrada pelo professor João Batista.

## Codificação das Máquinas de Turing na fita

Todas as máquinas que serão processadas pela *MTU* são codificadas da mesma maneira para padronizar o processamento. Aqui está um exemplo de uma *Máquina de Turing* com **2** estados, que incrementa infinitamente um número binário

> $00*00001*01011*0 1 0*11100*10010*1 011$  ^1000

Cada um dos símbolos tem um significado na fita a ser decodificada:

- $ -> Demarca início e fim das transições de estados da máquina a ser processada
- * -> Separa as transições presentes na máquina
- ^ -> Na seção fora da máquina, indica para onde o cabeçote aponta na fita de dados (à direita)

Além desses símbolos, há obviamente os símbolos 0, 1 e " " (vazio), que são lidos e escritos na seção de dados da fita.

### Codificação das transições

As transições na máquina sempre tem o mesmo formato, indepentende da quantidade de estados e transições que hajam. Elas sempre seguem esse mesmo formato:

> Estado Atual, Valor sendo lido, Próximo estado, Valor a ser escrito, Direção do cabeçote

Neste sentido, vamos pegar uma transição da fita acima, para exemplificar. A máquina acima possui dois estados. Da seguinte maneira:

> 11100

Para essa transição, lê-se:

- Estamos no estado **1**
- Lendo na fita o valor **1**
- Então vamos para o estado **1**
- Escrevemos **0** na fita
- E então movemos o cabeçote à **Esquerda** (0 = esquerda, 1 = direita)

Isso mantém-se independentemente do tamanho da máquina.

Os bits mais à esqurda da máquina, logo após o primeiro $, funcionam como um Program Counter (PC), onde guarda-se o estado atual e o valor atualmente lido na fita.

No exemplo da máquina citada acima, estamos no estado **0**, lendo **0** na fita. O número de bits à esquerda cresce dependendo da necessidade de bits para ditar os estados.

## Estados da máquina

A máquina foi implementada usando o software *JFLAP*, onde cada um dos estados da *MTU* é uma Máquina de Turing separada, para o processamento de cada etapa.

A máquina universal é uma máquina com seis etapas, seguindo a seguinte ordem:

- 1. Buscar o valor que o cabeçote está lendo e escrever na máquina;
- 2. Achar a instrução referente ao estado e valor atual;
- 3. Escrever na máquina o próximo estado e valor a ser escrito;
- 4. Escrever na fita de dados os valor salvo na máquina;
- 5. Mover o cabeçote para o valor indicado na transição;
- 6. Limpar a máquina para reverter as alterações feitas na fita de instruções.

A partir da execução desses passos, a *MTU* processa qualquer *Máquina de Turing* corretamente codificada com seção de dados.
