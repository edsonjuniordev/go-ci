# Continuous Integration (CI) com GitHub Actions  

Este projeto foi criado como parte de um estudo sobre **Continuous Integration (CI)** utilizando o GitHub Actions.  

## O que é CI?  
Continuous Integration, ou Integração Contínua, é uma prática de desenvolvimento de software que permite integrar alterações de código frequentemente e de forma automática em um repositório central. O objetivo é detectar erros mais cedo, garantindo a qualidade do código e a agilidade no desenvolvimento.  

## Objetivo deste projeto  
O principal objetivo deste projeto é:  
- Explorar os recursos do GitHub Actions para configurar pipelines de CI.  
- Aprender a automatizar tarefas, como execução de testes, linting e build do código.  
- Compreender o fluxo de integração e entrega contínua em projetos reais.  

## Como funciona o pipeline?  
1. O pipeline é disparado automaticamente em cada **pull request** no repositório.  
2. Ele realiza as seguintes etapas:  
   - **Setup da aplicação**: Instala as dependências e configura o ambiente necessário para execução.  
   - **Testes automatizados**: Executa os testes para validar o funcionamento do código.  
   - **Build da imagem Docker**: Cria uma imagem Docker da aplicação como parte do processo de CI.  

## Estrutura do Projeto  
- **.github/workflows**: Diretório onde estão configurados os arquivos YAML do GitHub Actions.  
- **Código de exemplo**: Código simples utilizado para validar os pipelines de CI configurados.  