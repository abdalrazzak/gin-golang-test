building docker images by following command

```
sudo docker-compose up --build 
```

##API Endpoint 
<table>
<tr>
<td> Method </td> <td> Endpoint </td><td> Request </td>
</tr>
<tr>
<td> POST </td>
<td> /login </td>
<td>

  
{
  
        "email" : "abboudbath4@gmai.com" ,
        "password" : "password" 
 
}
 

</td> 
</tr>

<tr>
<td> POST </td>
<td> /users </td>
<td>

  
{
 
        "age" :  ,
        "email" : "" ,
        "password" : "" 
 
}
 

</td> 
</tr>
<tr>
<td> GET </td>
<td> /users </td>
<td>
 
{
 
         
}
  
</td>
    
</tr>
    
<tr>
<td> GET </td>
<td> /users/{id} </td>
<td>
 body
 {   

}
      
</td>
    
</tr>
    
<tr>
<td> PUT </td>
<td> /users/{id} </td>
<td>
 body
 {  
        "age" :  , 
        "email" : "" ,
        "password" : "" 

}

header 
{  
      "Authorization" : token   
}

</td>
    
</tr>

<tr>
<td> DELETE </td>
<td> /users/{id} </td>
<td>
 body
 {   
}

header 
{  
      "Authorization" : token   
}

</td>
    
</tr>

<tr>
<td> POST</td>
<td> /files </td>
<td>
 body
 {   
    "content" : "djadsajd.jpg"
}

header 
{  
      "Authorization" : token   
}

</td>
    
</tr>

<tr>
<td> Get</td>
<td> /files </td>
<td>
 body
 {   
    
}

header 
{  
      "Authorization" : token   
}

</td>
    
</tr>

<tr>
<td> DELETE</td>
<td> /files/{id} </td>
<td>
 body
 {   
    
}

header 
{  
      "Authorization" : token   
}

</td>
    
</tr>
 
</table>






For Running Tests

```
sudo docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit
```