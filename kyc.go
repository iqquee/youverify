package youverify

type (
	// BankVerificationNumberRequest is the request object for BankVerificationNumber() to verify a BVN
	BankVerificationNumberRequest struct {
		BVN              string                 `json:"id"`
		MetaData         map[string]interface{} `json:"metadata"`
		IsSubjectConsent bool                   `json:"isSubjectConsent"` // This field must be true
		PremiumBVN       bool                   `json:"premiumBVN"`
	}

	// BankVerificationNumberResponse is the response object for the BankVerificationNumber()
	BankVerificationNumberResponse struct {
		Success    bool   `json:"success"`
		StatusCode int    `json:"statusCode"`
		Message    string `json:"message"`
		Data       struct {
			Id                    string      `json:"id"`
			ParentId              interface{} `json:"parentId"`
			Status                string      `json:"status"`
			Reason                interface{} `json:"reason"`
			DataValidation        bool        `json:"dataValidation"`
			SelfieValidation      bool        `json:"selfieValidation"`
			FirstName             string      `json:"firstName"`
			MiddleName            interface{} `json:"middleName"`
			LastName              string      `json:"lastName"`
			Image                 string      `json:"image"`
			EnrollmentBranch      interface{} `json:"enrollmentBranch"`
			EnrollmentInstitution interface{} `json:"enrollmentInstitution"`
			Mobile                string      `json:"mobile"`
			DateOfBirth           string      `json:"dateOfBirth"`
			IsConsent             bool        `json:"isConsent"`
			IdNumber              string      `json:"idNumber"`
			Nin                   string      `json:"nin"`
			ShouldRetrivedNin     bool        `json:"shouldRetrivedNin"`
			BusinessId            string      `json:"businessId"`
			Type                  string      `json:"type"`
			AllValidationPassed   bool        `json:"allValidationPassed"`
			RequestedAt           string      `json:"requestedAt"`
			RequestedById         string      `json:"requestedById"`
			Country               string      `json:"country"`
			CreatedAt             string      `json:"createdAt"`
			LastModifiedAt        string      `json:"lastModifiedAt"`
			Email                 interface{} `json:"email"`
			RegistrationDate      string      `json:"registrationDate"`
			Gender                string      `json:"gender"`
			LevelOfAccount        interface{} `json:"levelOfAccount"`
			Address               struct {
				Town        interface{} `json:"town"`
				Lga         interface{} `json:"lga"`
				State       interface{} `json:"state"`
				AddressLine interface{} `json:"addressLine"`
			} `json:"address"`
			Title         interface{} `json:"title"`
			MaritalStatus interface{} `json:"maritalStatus"`
			LgaOfOrigin   interface{} `json:"lgaOfOrigin"`
			OtherMobile   interface{} `json:"otherMobile"`
			StateOfOrigin interface{} `json:"stateOfOrigin"`
			WatchListed   interface{} `json:"watchListed"`
			NameOnCard    interface{} `json:"nameOnCard"`
			FullDetails   bool        `json:"fullDetails"`
			Metadata      interface{} `json:"metadata"`
			RequestedBy   struct {
				FirstName  string `json:"firstName"`
				LastName   string `json:"lastName"`
				MiddleName string `json:"middleName"`
				Id         string `json:"id"`
			} `json:"requestedBy"`
		} `json:"data"`
		Links []interface{} `json:"links"`
	}
)

// BankVerificationNumber is method used to verify BVN(Bank Verification Number) in Nigeria
func (n *Nigeria) BankVerificationNumber(request BankVerificationNumberRequest) (*BankVerificationNumberResponse, error) {
	if !request.IsSubjectConsent {
		return nil, errSubjectConsent
	}

	url := "identity/ng/bvn"
	var response BankVerificationNumberResponse
	if err := newRequest(methodPOST, url, request, response); err != nil {
		return nil, err
	}

	return &response, nil
}
