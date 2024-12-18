import React, { useState } from 'react';
import './Body.css';
import axios from 'axios';

const Body = () => {
  const [token, setToken] = useState('');
  const [amount, setAmount] = useState(0);
  const [priceData, setPriceData] = useState([]); 
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(false);


  const getTokenValue = async (token, amount) => {
    setLoading(true);
    setError(null); 

    try {
      
      const response = await axios.get('http://localhost:8080/tokenprice', {
        params: {
          id: token,
          amount: amount,
        },
      });

      console.log(response.data);
      setPriceData(response.data.data || []); 
    } catch (err) {
      console.error(err);
      setError('Failed to fetch price data.');
    } finally {
      setLoading(false);
    }
  };

  
  const handleSubmit = (e) => {
    e.preventDefault(); 

    if (!token || amount <= 0) {
      setError('Please enter a valid token and amount.');
      return;
    }

    getTokenValue(token, amount); 
  };

  return (
    <div className='container'>
      <h2>Welcome to Coin Pricer</h2>
      <p>
        Looking to find the latest prices for your favorite cryptocurrencies? <br />
        Coin Pricer makes it easy to search and track real-time prices for thousands of crypto tokens. 💰🪙
      </p>

      
      <form onSubmit={handleSubmit}>
        <input
          placeholder='Crypto Token'
          type='text'
          value={token}
          onChange={(e) => setToken(e.target.value)}
        />
        <input
          placeholder='Amount'
          type='number'
          value={amount}
          onChange={(e) => setAmount(parseFloat(e.target.value))} // Ensure amount is a number
        />
        <button type='submit'>Submit</button>
      </form>

   
      {loading && <p>Loading...</p>}

     
      {error && <p style={{ color: 'red' }}>{error}</p>}

    
      {priceData.length > 0 && (
        <div>
          <h3>Price Data</h3>
          <ul>
            {priceData.map((token) => (
              <li key={token.id}>
                <h4>{token.name} ({token.symbol})</h4>
                <p>Amount: {token.amount}</p>
                <p>Price (USD): ${token.quote.USD.price.toFixed(6)}</p>
                <p>Last Updated: {new Date(token.last_updated).toLocaleString()}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default Body;
