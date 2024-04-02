import React from 'react';
import './App.css'; // Make sure this line is there to import your styles
import TimeSeriesChart from './TimeSeriesChart';

function App() {
  return (
    <div className="App">
      <div className="chart-container">
        <div className="chart-box">
          <TimeSeriesChart url="http://localhost:8080/api/series1" title="Sine Wave" />
        </div>
        <div className="chart-box">
          <TimeSeriesChart url="http://localhost:8080/api/series2" title="Damped Oscillations" />
        </div>
      </div>
    </div>
  );
}

export default App;
