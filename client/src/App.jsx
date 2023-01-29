import React from 'react'

function App() {
  return (
    <div>
      <h1>Hello world</h1>
      <button onClick={async () => {
        const response = await fetch('http://localhost:3000/users')
        const data = await response.json()
        console.log(data)
      }}>Optain data</button>
    </div>
  )
}

export default App    