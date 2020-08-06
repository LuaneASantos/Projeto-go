# Projeto-go

Neste projeto consumi a api https://superheroapi.com, inseri no banco PostGreSQl os dados, e manipulei os mesmos.

<h2> API's </h2> 

<table>
   <thead>
      <th>Patch</th>
      <th>Método</th>
      <th>Descrição</th>
   </thead>
   <tbody>
      <tr>
         <td>/all </td>
         <td>GET</td>
         <td>Retorna todos os registros.</td>
      </tr>
      <tr>
         <td>/heroes </td>
         <td>GET</td>
         <td>Retorna todos os super-heróis.</td>
      </tr>
      <tr>
         <td>/villains </td>
         <td>GET</td>
         <td>Retorna todos os vilões.</td>
      </tr>
      <tr>
         <td>/search/{name} </td>
         <td>GET</td>
         <td>Retorna apenas o super do nome selecionado.</td>
      </tr>
      <tr>
         <td>/id/{id} </td>
         <td>GET</td>
         <td>Retorna apenas o registro do id selecionado.</td>
      </tr>
      <tr>
         <td>/id/{id} </td>
         <td>DELETE</td>
         <td>Deleta apenas o registro do id selecionado.</td>
      </tr>
      <tr>
         <td>/new </td>
         <td>POST</td>
         <td>Cria um novo registro.</td>
      </tr>
   </tbody>
</table>

<h2> Versão GoLang </h2> 

<a href="https://golang.org/doc/install?download=go1.14.6.windows-amd64.msi">Go 1.14.6</a>

<h2> Bibliotecas e Ferramentas </h2> 

gorilla/mux <p>
lib/pq <p>
database/sql <p>
encoding/json <p>
fmt <p>
net/http <p>
net/http/httptest <p>
io/ioutil <p>

<h2> Como rodar aplicação</h2> 

<h3>db.go</h3> 

Neste arquivo é criado o banco de dados e inserido os registros, então é necessario rodar o mesmo com o comando "go run db.go"

<h3>main.go</h3> 

Neste arquivo é instanciando as rotas das api's, para compila-lo devesse rodar o comando  "go run main.go"

<h3>main_test.go</h3> 

Neste arquivo é feito os testes das api's, para rodar os testes devesse digitar o comando  "go test -v"
