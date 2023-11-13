import './App.css';
import { useEffect, useState } from 'react';
import { Button, Container } from 'semantic-ui-react';
import { checkForExistingVisitor, getCurrentNumber, getNextNumber, getPreviousNumber } from './api/api';

function App() {
  const [currentValue, setCurrentValue] = useState("Press current to start");

  useEffect(() => {
    checkForExistingVisitor();
  }, []);

  const getValue = async (fn) => {
    try {
      const value = await fn();

      setCurrentValue(value)
    }
    catch(exception){
      alert("API Error");
    }    
  }
  
  return (
    <div className="App">
      <Container className='mainContainer'>
        <h2>{currentValue}</h2>
        <Button onClick={() => getValue(getPreviousNumber)}>Previous</Button>
        <Button onClick={() => getValue(getCurrentNumber)}>Current</Button>
        <Button onClick={() => getValue(getNextNumber)}>Next</Button>      
      </Container>      
    </div>
  );
}

export default App;
