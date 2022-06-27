# BasicCarTrackNetwork

1) Generate Crypto-Materials through create-artifacts.sh
2) Network Service Up By Docker-compose up -d command in docker-compose directory
3) Create channel through createchannel.sh script
4) Deploy Chaincode in network through DeployChaincodeContractApi.sh Script
 **Invoke and query smart contract through DeployChaincodeContractApi.sh script (Through Commented portion code for different invoke function and query)**

5) npm install in api-server directory
6) Through command **node app.js** up the server
7) Using Postman
8) Regitster the User By Post Method: input Json Format: EX: {"userId":"samyakjain","orgMSP":"Org2MSP" }
9) Invoke using Tx APi:EX: { "userId":"samyakjain", "orgMSP":"Org2MSP", "channelName":"mychannel", "chaincodeName":"TrackCar", "data {"function":"ManufactureCar","carNumber":"1001","make":"Ola","model":"v1","colour":"white","owner":"manufacturer","state":"Created"}
10) Similarly All Invoke Functions 
11)For Query Use /query api and similar json format.
    
    
    
          
