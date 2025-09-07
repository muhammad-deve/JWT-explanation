# ✨ JWT Token (Step by Step Notes) ✨

---

## 🔑 Token Types

### 🟢 Access Token
- Short-lived (**15 min – 1 hour**)  
- Used to access protected APIs  
- Must be refreshed when expired  

### 🔄 Refresh Token
- Long-lived (**days or weeks**)  
- Used only to request new Access Tokens  
- If expired → user must log in again  

---

## 🔐 Authentication Flow

### 📝 Sign Up / Sign In
- Backend creates **Access Token + Refresh Token**  
- Sends both tokens to frontend (often stored in cookies)  

### 🌍 Accessing APIs
- Frontend sends **Access Token** with requests  
- If valid → ✅ request succeeds  
- If expired → ⏳ frontend calls **Refresh API**  
  > *(Usually, the Access Token does not actually “expire” for the user because the frontend proactively calls the Refresh Token API to renew it before expiration.)*  

### 🔄 Refreshing Tokens
- `POST /user/refresh-token`  
- Frontend sends **Refresh Token**  
- Backend checks validity  

**Results:**  
- ✅ If valid → returns **new Access Token** (sometimes also a new Refresh Token)  
  > *(I prefer updating the Refresh Token as well)*  
- ❌ If expired/invalid → user must log in again  

💡 **Note:** Both tokens will never expire if you update the Refresh Token every time you update the Access Token.  

---

## 🌐 API Endpoints

```http
POST /user/sign-up          → returns Access + Refresh Tokens
POST /user/sign-in          → returns Access + Refresh Tokens
POST /user/refresh-token    → returns new Access Token (and maybe new Refresh Token) <------- LIKE THIS
