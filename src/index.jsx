import React from 'react'
import ReactDOM from 'react-dom/client'
import './styles/main.css'
import Home from './pages/Home'

const App = () => {
  return (
    <div className="app">
      <Home />
    </div>
  )
}

const root = ReactDOM.createRoot(document.getElementById('root'))
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
)
