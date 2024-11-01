import React from 'react'
import './Body.css'

const Body = () => {
  return (
        <div className='container'>
        <h2>Welcome to Coin Pricer</h2>
        <p>Looking to find the latest prices for your favorite cryptocurrencies? <br/>Coin Pricer makes it easy to search and track real-time prices for thousands of crypto tokens. 💰🪙</p>
        <input placeholder='Crypto Token' type='text'/>
        <input placeholder='Amount' type='number'/>
      </div>
  )
}

export default Body