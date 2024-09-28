import './App.css'
import { Layout } from './comp/layout'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import { Addhabit } from './comp/addhabit'

function App() {

  return (
    <Router>
    <Routes>
        <Route path="/" element={<Layout />} />
        <Route path="/add-habit" element={<Addhabit />} />
        
    </Routes>
</Router>
);
};



export default App
