# Golang_Secure_Redis_Communication
![Adsız tasarım](https://user-images.githubusercontent.com/92402372/205440598-5f3fa57d-da8b-4dc0-8ec8-e96b47a80da9.jpg)

1- We send encrypted data to the relevant redis key according to the operation we want to do from the client.
2- The server is constantly listening to all defined redis keys, decrypting the incoming data, processing it and returning the relevant result.
