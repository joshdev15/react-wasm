import React from 'react'
import ReactDOM from 'react-dom/client'
import Home from './pages/Home'
import './styles/main.css'

const root = ReactDOM.createRoot(document.getElementById('root'))
root.render(
  <React.StrictMode>
    <div className="app">
      <Home />
    </div>
  </React.StrictMode>
)
