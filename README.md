<div align="center">
<h1>GOLANG</h1>
<p>
<img src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/1200px-Go_Logo_Blue.svg.png" alt="Golang" width=200px/>
</p>
<br>
<p > <b>The Golang project is an API to create posts in a simple social media<b> </p>
 
 ![GO](https://camo.githubusercontent.com/61396e7d6de2342ea69b46d8d1df6c939b7bfad3f63fe934762da9051e734d66/68747470733a2f2f696d672e736869656c64732e696f2f7374617469632f76313f7374796c653d666f722d7468652d6261646765266d6573736167653d476f26636f6c6f723d303041444438266c6f676f3d476f266c6f676f436f6c6f723d464646464646266c6162656c3d)
 </div>
 
 ## INSTALLATION
 
 ```
# Clone this repository
$ git clone https://github.com/Rodolpho-Oliveira/projeto-golang.git

# Go into the repository
$ cd projeto-golang

# Install dependencies
$ go get

 ```
 
 ## USAGE
 
 POST ```/sign-up```<br>
 
 Need to receive through the body a parameter ```username``` and an ```avatar```
 ```
 {
    username: "randomName",
    avatar: "https://www...."
 }
 ```
 
 POST ```/tweets```<br>
 
 Need to receive through the body a parameter ```username``` and ```tweet```
 ```
 {
    username: "randomName",
    tweet: "I want to hire this dev!"
 }
 ```
 
 GET ```/tweets```<br>
 
 Return the last 10 tweets
 ```
 [
   {
      username: "randomName",
      avatar: "https://www...",
      tweet: "I want to hire this dev!"
   }
 ]
 ```
 
 
