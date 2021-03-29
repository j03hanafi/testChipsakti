package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Found error in log ", err)
	}
	log.SetOutput(file)

	fmt.Println("Test Started!")

	start := time.Now()
	defer func() {
		log.Printf("\nExecution Time: %v\n\n", time.Since(start))
	}()

	customer_no := []string{"1", "2", "3", "4", "5"}

	var wg sync.WaitGroup

	for _, cusNo := range customer_no {
		wg.Add(5)

		go reqPPOBInquiry(cusNo, &wg)
		log.Printf("PPOBInquiry %v is running now\n", cusNo)

		go reqPPOBPayment(cusNo, &wg)
		log.Printf("PPOBPayment %v is running now\n", cusNo)

		go reqPPPOBStatus(cusNo, &wg)
		log.Printf("PPOBStatus %v is running now\n", cusNo)

		go reqTopupBuy(cusNo, &wg)
		log.Printf("TopupBuy %v is running now\n", cusNo)

		go reqTopupCheck(cusNo, &wg)
		log.Printf("TopupCheck %v is running now\n", cusNo)

		log.Println("")
	}
	log.Println("")

	wg.Wait()
	fmt.Println("Test Finished!")
	//reqChipsakti(customer_no[1])
}

func reqPPOBInquiry(customer_no string, wg *sync.WaitGroup) PPOBInquiryResponse {
	var response PPOBInquiryResponse
	var baseUrl = "http://localhost:6010/ppob/inquiry"

	reqBody, _ := json.Marshal(PPOBInquiryRequest{
		TransactionID: "2021",
		PartnerID:     "USER01",
		ProductCode:   "WOM",
		CustomerNo:    customer_no,
		Periode:       "2020",
		MerchantCode:  "KIOS01",
		RequestTime:   "2018-05-15 15:10:05",
	})

	responseBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(baseUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	log.Printf("PPOBInquiry %v || %+v\n", customer_no, response)
	wg.Done()
	return response
}

func reqPPOBPayment(customer_no string, wg *sync.WaitGroup) PPOBPaymentResponse {
	var response PPOBPaymentResponse
	var baseUrl = "http://localhost:6010/ppob/payment"

	reqBody, _ := json.Marshal(PPOBPaymentRequest{
		TransactionID: "2015",
		PartnerID:     "USER01",
		ProductCode:   "WOM",
		CustomerNo:    customer_no,
		MerchantCode:  "KIOS01",
		ReffID:        "12345",
		Amount:        873300,
		RequestTime:   "2018-05-15 15:10:05",
	})

	responseBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(baseUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	log.Printf("PPOBPayment %v || %+v\n", customer_no, response)
	wg.Done()
	return response
}

func reqPPPOBStatus(customer_no string, wg *sync.WaitGroup) PPOBStatusResponse {
	var response PPOBStatusResponse
	var baseUrl = "http://localhost:6010/ppob/status"

	reqBody, _ := json.Marshal(PPOBStatusRequest{
		TransactionID: "2021",
		PartnerID:     "USER01",
		ProductCode:   "WOM",
		CustomerNo:    customer_no,
		MerchantCode:  "KIOS01",
		ReffID:        "12345",
		Amount:        10000,
		RequestTime:   "2018-05-15 15:10:05",
	})

	responseBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(baseUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	log.Printf("PPOBStatus %v || %+v\n", customer_no, response)
	wg.Done()
	return response
}

func reqTopupBuy(customer_no string, wg *sync.WaitGroup) TopupBuyResponse {
	var response TopupBuyResponse
	var baseUrl = "http://localhost:6010/topup/buy"

	reqBody, _ := json.Marshal(TopupBuyRequest{
		TransactionID: "2021",
		PartnerID:     "USER01",
		ProductCode:   "WOM",
		CustomerNo:    customer_no,
		MerchantCode:  "KIOS01",
		RequestTime:   "2018-05-15 15:10:05",
	})

	responseBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(baseUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	log.Printf("TopupBuy %v || %+v\n", customer_no, response)
	wg.Done()
	return response
}

func reqTopupCheck(customer_no string, wg *sync.WaitGroup) TopupCheckResponse {
	var response TopupCheckResponse
	var baseUrl = "http://localhost:6010/topup/check"

	reqBody, _ := json.Marshal(TopupCheckRequest{
		TransactionID: "2021",
		PartnerID:     "USER01",
		ProductCode:   "WOM",
		CustomerNo:    customer_no,
		MerchantCode:  "KIOS01",
		RequestTime:   "2018-05-15 15:10:05",
	})

	responseBody := bytes.NewBuffer(reqBody)
	resp, err := http.Post(baseUrl, "application/json", responseBody)

	if err != nil {
		log.Fatal("An error occured", err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &response)

	log.Printf("TopupCheck %v || %+v\n", customer_no, response)
	wg.Done()
	return response
}
